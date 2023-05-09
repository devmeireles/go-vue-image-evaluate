import axios from "axios";
import { defineStore } from "pinia";

export interface Report {
  id: string;
  name: string;
}

interface ReportState {
  reports: Report[];
}

export const useReportStore = defineStore({
  id: "report",
  state: () =>
    ({
      reports: []
    } as ReportState),
  getters: {
    getReportById: (state) => (id: string) => {
      return state.reports.find((reports: Report) => reports.id === id);
    },
    getReports: (state) => () => state.reports
  },
  actions: {
    async fetchReports() {
      try {
        const response = await axios.get("http://localhost:3000/report");

        const { data, status } = response;

        if (data.success && status === 200) {
          this.$patch({
            reports: data.data
          });
        }
      } catch (error) {
        console.log(error);
      }
    },

    async saveReport(payload: Record<string, any>) {
      try {
        const response = await axios.post("http://localhost:3000/report", {
          external_id: payload.externalID,
          image_url: payload.imageURL
        });

        const { data, status } = response;

        console.log(response);
      } catch (error) {
        console.log(error);
      }
    }
  }
});
