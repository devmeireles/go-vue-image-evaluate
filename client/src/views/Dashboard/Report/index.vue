<template>
  <CustomCard :title="$t('report.title_plural')">
    <ReportsList />
  </CustomCard>
</template>

<script lang="ts">
import CustomCard from "@/components/organisms/CustomCard.vue";
import ReportsList from "@/components/organisms/ReportsList.vue";
import { defineComponent } from "vue";

export default defineComponent({
  name: "ListReport",
  components: { CustomCard, ReportsList }
});
</script>
