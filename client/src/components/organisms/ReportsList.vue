<template>
  <v-table>
    <thead>
      <tr>
        <th class="text-left">
          #
        </th>
        <th class="text-left">
          {{ $t('report.image_url') }}
        </th>
        <th class="text-left">
          {{ $t('report.status') }}
        </th>
        <th class="text-left">
          {{ $t('report.priority') }}
        </th>
        <th class="text-left">
          {{ $t('report.actions') }}
        </th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="item in data"
        :key="item.name"
      >
        <td>{{ item.external_id }}</td>
        <td class="image-url">{{ item.image_url }}</td>
        <td><report-status-chip item-type="status" :status="item.status"/></td>
        <td><report-status-chip item-type="priority" :status="item.priority"/></td>
        <td>
          <v-btn
            color="primary"
            variant="outlined"
            class="cta-btn"
            @click="() => $router.push({
              name: 'report-evaluate',
              params: {
                id: item.id
              }
            })"
          >
            {{ $t('actions.evaluate') }}
            <v-icon
              end
              icon="mdi-open-in-new"
            ></v-icon>
          </v-btn>
        </td>
      </tr>
    </tbody>
  </v-table>
</template>

<script lang="ts">
import ReportStatusChip from '@/components/atoms/ReportStatusChip.vue';
import { routes } from '@/consts/routes';

export default {
  name: "reports-list",
  data: () => ({
    evaluateRoute: routes.report.evaluate
  }),
  props: {
    data: {
      type: Object,
      default: () => { }
    }
  },
  components: { ReportStatusChip }
}
</script>

<style lang="scss" scoped>
.image-url {
  max-width: 150px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.cta-btn {
  text-transform: initial;
}
</style>