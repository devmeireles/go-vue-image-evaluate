const MAIN_ROUTE = `/dashboard`;

export const routes = {
  dashboard: {
    main: `${MAIN_ROUTE}`
  },
  report: {
    main: `${MAIN_ROUTE}/report`,
    create: `${MAIN_ROUTE}/report/create`,
    evaluate: `${MAIN_ROUTE}/report/evaluate`
  }
};
