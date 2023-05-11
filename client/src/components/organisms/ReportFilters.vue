<template>
  <v-container>
    <v-row>
      <v-col cols="6" md="6">
        <v-select
          :label="$t('report.status')"
          v-model="formData.status"
          :items="status"
          item-title="label"
        ></v-select>
      </v-col>
      
      <v-col cols="6" md="6">
        <v-select
          :label="$t('report.priority')"
          v-model="formData.priority"
          :items="priorities"
          item-title="label"
        ></v-select>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
export default ({
  name: 'report-filters',
  data: () => {
    return {
      formData: {
        priority: 0,
        status: 1
      },
      priorities: [],
      status: []
    }
  },
  watch: {
    'formData.priority': function (value) {
      this.$emit('updateFilter', {
        input: 'priority',
        value
      });
    },

    'formData.status': function (value) {
      this.$emit('updateFilter', {
        input: 'status',
        value
      });
    }
  },

  mounted() {
    const priorities = [
      {
        value: 0,
        label: this.$t('report.all')
      },
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
        value: 0,
        label: this.$t('report.all')
      },
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
  }
})
</script>