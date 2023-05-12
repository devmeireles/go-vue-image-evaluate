<template>
  <CustomCard :title="$t('core.report')" :is-loading="isLoading">
    <v-form ref="form">
      <v-container>
        <v-row>
          <v-col cols="6" md="6">
            <v-text-field
              v-model="formData.external_id"
              :label="$t('report.external_id')"
              :error-messages="v$.formData.external_id.$errors.map(e => e.$message)"
              variant="outlined"
              shaped
              :readonly="editingMode"
              :disabled="editingMode"
              required
            ></v-text-field>
          </v-col>

          <v-col cols="6" md="6">
            <v-text-field
              v-model="formData.image_url"
              :label="$t('report.image_url')"
              :error-messages="v$.formData.image_url.$errors.map(e => e.$message)"
              variant="outlined"
              shaped
              :readonly="editingMode"
              :disabled="editingMode"
              required
            ></v-text-field>
            <v-btn v-if="editingMode" color="primary" class="btn-image" variant="plain" :onclick="() => openImage(formData.image_url)">{{ $t('report.open_image') }}</v-btn>
          </v-col>
        </v-row>

        <v-row v-if="editingMode">
          <v-col cols="6" md="6">
            <v-select
              :label="$t('report.priority')"
              v-model="formData.priority"
              :items="priorities"
              item-title="label"
            ></v-select>
          </v-col>

          <v-col cols="6" md="6">
            <v-select
              :label="$t('report.status')"
              v-model="formData.status"
              :items="status"
              item-title="label"
            ></v-select>
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12" md="12">
            <div class="d-flex flex-row justify-end">
              <v-btn
                :loading="isLoading"
                :disabled="isLoading"
                color="warning"
                variant="outlined"
                width="200"
                class="mt-6 mr-5"
                @click="cancelAction()"
              >
                {{ $t('actions.close') }}
              </v-btn>
              <v-btn
                :loading="isLoading"
                :disabled="isLoading"
                color="success"
                variant="outlined"
                width="200"
                class="mt-6"
                @click="submitForm"
              >
                {{ editingMode ? $t('actions.update') : $t('actions.save') }}
              </v-btn>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-form>
  </CustomCard>
</template>

<script lang="ts">
import CustomCard from "@/components/organisms/CustomCard.vue";
import { routes } from "@/consts/routes";
import { useCoreStore } from "@/stores/core";
import { useReportStore } from "@/stores/report";
import type { CreateReportDTO } from "@/types/TReport";
import useVuelidate from "@vuelidate/core";
import { required } from "@vuelidate/validators";
import { defineComponent } from "vue";

export default defineComponent({
  name: "FormReport",
  data: () => {
    const formData: CreateReportDTO = {
      external_id: "",
      image_url: "",
      status: undefined,
      priority: undefined,
    };

    return {
      v$: useVuelidate(),
      isLoading: false,
      editingMode: false,
      priorities: [],
      status: [],
      currentID: '',
      formData,
    }
  },
  components: { CustomCard },
  async mounted() {
    if (this.$route.params.id) {
      this.isLoading = true;
      this.currentID = this.$route.params.id as string;
      this.editingMode = true;

      const priorities = [
        {
          value: 1,
          label: this.$t('report.low')
        },
        {
          value: 2,
          label: this.$t('report.medium')
        },
        {
          value: 3,
          label: this.$t('report.high')
        }
      ];

      const status = [
        {
          value: 1,
          label: this.$t('report.created')
        },
        {
          value: 2,
          label: this.$t('report.reviewing')
        },
        {
          value: 3,
          label: this.$t('report.reviewed')
        }
      ];

      this.priorities = priorities;
      this.status = status;

      const reportStore = useReportStore();
      const data = await reportStore.fetchReport(this.currentID);
      this.isLoading = false;

      Object.assign(this.formData, { ...data });
    }
  },
  methods: {
    cancelAction() {
      this.$router.push(routes.report.main);
    },
    async submitForm() {
      this.isLoading = true;
      const reportStore = useReportStore();
      const coreStore = useCoreStore();
      await this.v$.$validate();

      if (!this.v$.$error) {
        try {

          let data;

          if (this.editingMode) {
            data = await reportStore.patchReport(this.formData);
            coreStore.openSnackbar(`${data.external_id} updated sucessfuly`)
          } else {
            data = await reportStore.saveReport(this.formData);
            coreStore.openSnackbar(`${data.external_id} created sucessfuly`)
          }

          this.$router.push(routes.report.main);

        } catch (error) {
          console.log(error)
        }
      }

      this.isLoading = false;
    },
    openImage(imageURL: string) {
      window.open(imageURL, "_blank");
    }
  },
  validations() {
    return {
      formData: {
        external_id: { required },
        image_url: { required },
      }
    }
  }
});
</script>

<style lang="scss" scoped>
.btn-image {
  color: var(--primary);
  cursor: pointer;
}
</style>