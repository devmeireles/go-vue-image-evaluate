<template>
  <CustomCard>
    <v-form v-model="valid">
      <v-container>
        <v-row no-gutters>
          <v-col cols="12" md="12">
            <v-text-field
              v-model="firstname"
              :rules="nameRules"
              :counter="10"
              label="First name"
              variant="outlined"
              shaped
              required
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row no-gutters>
          <v-col cols="12" md="12">
            <v-textarea
              outlined
              rows="4"
              hint="help text"
              name="input-7-4"
              label="Description"
            ></v-textarea>
          </v-col>
        </v-row>

        <v-row no-gutters class="d-flex flex-row justify-space-around">
          <v-col cols="12" md="8">
            <v-file-input
              chips
              label="File input w/ chips"
              prepend-icon=""
              append-icon="mdi-camera"
            ></v-file-input>
          </v-col>

          <v-col cols="12" md="4">
            <div class="d-flex flex-row align-center justify-end">
              <div class="pa-2 mx-5">
                <p>Status</p>
              </div>
              <div>
                <v-btn-toggle v-model="toggle_one" tile mandatory>
                  <v-btn
                    value="left"
                    color="success"
                    size="small"
                    variant="tonal"
                  >
                    On
                  </v-btn>
                  <v-btn
                    value="center"
                    color="error"
                    size="small"
                    variant="tonal"
                  >
                    Off
                  </v-btn>
                </v-btn-toggle>
              </div>
            </div>
          </v-col>
        </v-row>

        <v-row no-gutters>
          <v-col cols="12" md="12">
            <div class="d-flex flex-row justify-end">
              <v-btn
                class="ma-2"
                :loading="loading"
                :disabled="loading"
                color="success"
                variant="outlined"
                width="150"
                @click="loader = 'loading'"
              >
                Save
              </v-btn>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-form>
  </CustomCard>
</template>

<script lang="ts">
  import { defineComponent } from "vue";
  import CustomCard from "@/components/organisms/CustomCard.vue";

  export default defineComponent({
    name: "CreateReport",
    data: () => ({
      loader: null,
      loading: false,
      toggle_one: "left",
      valid: false,
      firstname: "",
      lastname: "",
      nameRules: [
        (v) => !!v || "Name is required",
        (v) => v.length <= 10 || "Name must be less than 10 characters"
      ],
      email: "",
      emailRules: [
        (v) => !!v || "E-mail is required",
        (v) => /.+@.+/.test(v) || "E-mail must be valid"
      ]
    }),
    watch: {
      loader() {
        const l = this.loader;
        this[l] = !this[l];
        setTimeout(() => (this[l] = false), 3000);
        this.loader = null;
      }
    },
    components: { CustomCard }
  });
</script>
