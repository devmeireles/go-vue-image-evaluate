<template>
  <v-chip variant="text" :text="renderText" />
</template>

<script lang="ts">
import { PRIORITY, STATUS } from '@/consts/enums';
import { defineComponent } from 'vue';

export default defineComponent({
  name: "report-status-chip",
  props: {
    status: {
      type: Number,
      default: 1,
    },
    itemType: {
      type: String, // as status or priority
    }
  },
  computed: {
    renderText() {
      let res = '';
      let translation = '';

      if (this.itemType === 'status') {
        translation = STATUS[this.status];
      } else {
        translation = PRIORITY[this.status];
      }

      return this.$t(`report.${translation}`);
    }
  }
});
</script>