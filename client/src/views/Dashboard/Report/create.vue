<template>
  <CustomCard :title="$t('core.report')">
    <v-form ref="form">
      <v-container>
        <v-row>
          <v-col cols="6" md="6">
            <v-text-field
              v-model="formData.externalID"
              :label="$t('report.external_id')"
              :error-messages="v$.formData.externalID.$errors.map(e => e.$message)"
              variant="outlined"
              shaped
              required
            ></v-text-field>
          </v-col>

          <v-col cols="6" md="6">
            <v-text-field
              v-model="formData.imageURL"
              :label="$t('report.image_url')"
              :error-messages="v$.formData.imageURL.$errors.map(e => e.$message)"
              variant="outlined"
              shaped
              required
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12" md="12">
            <div class="d-flex flex-row justify-end">
              <v-btn
                :loading="isLoading"
                :disabled="isLoading"
                color="success"
                variant="outlined"
                width="200"
                class="mt-6"
                @click="submitForm"
              >
                {{ $t('actions.save') }}
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
import { useReportStore, type Report } from "@/stores/report";
import useVuelidate from "@vuelidate/core";
import { required } from "@vuelidate/validators";
import { defineComponent } from "vue";

export default defineComponent({
  name: "CreateReport",
  data: () => ({
    v$: useVuelidate(),
    isLoading: false,
    formData: {
      externalID: "",
      imageURL: "",
    }
  }),
  components: { CustomCard },
  methods: {
    async submitForm() {
      const reportStore = useReportStore();

      this.isLoading = true;
      await this.v$.$validate();

      if (!this.v$.$error) {
        const data = this.formData as unknown as Report;
        reportStore.saveReport(data);
      }


      this.isLoading = false;
    }
  },
  validations() {
    return {
      formData: {
        externalID: { required },
        imageURL: { required },
      }
    }
  }
});
</script>
