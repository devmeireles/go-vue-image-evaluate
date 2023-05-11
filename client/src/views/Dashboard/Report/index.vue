<template>
  <CustomCard :title="$t('report.title_plural')" :createRoute="createRoute">
    <ReportFilters @update-filter="updateFilter" class="mb-8"/>
    <ReportsList :data="reports" />
  </CustomCard>
</template>

<script lang="ts" setup>
import CustomCard from "@/components/organisms/CustomCard.vue";
import ReportFilters from "@/components/organisms/ReportFilters.vue";
import ReportsList from "@/components/organisms/ReportsList.vue";
import { routes } from "@/consts/routes";
import { useReportStore } from "@/stores/report";
import { computed, onMounted } from "vue";

const createRoute = routes.report.create;

const reportStore = useReportStore();

const filters: { [x: string]: number; }[] = [];

const updateFilter = (event: { input: string, value: number }) => {
  filters.push({ [event.input]: event.value })
  reportStore.fetchReports(filters);
}

onMounted(() => {
  reportStore.fetchReports();
});

const reports = computed(() => {
  return reportStore.getReports();
});
</script>
