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
        <td><v-chip color="warning" :text="$t('report.waiting')"></v-chip></td>
        <td><v-chip color="warning" :text="$t('report.waiting')"></v-chip></td>
        <td>
          <v-btn
            color="primary"
            variant="outlined"
            class="cta-btn"
            @click=""
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
export default {
  name: 'reports-list',
  props: {
    data: {
      type: Object,
      default: () => { }
    }
  },
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