<!-- eslint-disable vue/no-mutating-props -->
<script setup lang="ts" generic="T">
import type { PresetDate } from 'ant-design-vue/es/vc-picker/interface';
import type { Dayjs } from 'dayjs';

import { computed, type MaybeRef, ref } from 'vue';

import { Add, Search } from '@vben/icons';
import { $t } from '@vben/locales';
import { VbenIcon } from '@vben-core/shadcn-ui';

import {
  Badge,
  Button,
  Col,
  Flex,
  Input,
  Popover,
  RangePicker,
  Row,
  Select,
  SelectOption,
} from 'ant-design-vue';

import { dateUtil, disableFutureDates } from '#/utils/helper/date-util';
import {
  FieldType,
  getGetter,
  NumberOp,
  SelectOp,
  StringOp,
  useQueryForm,
} from '#/utils/pagi/form';
import { Operation, type QueryPagi } from '#/utils/pagi/query';

defineOptions({ name: 'Filter' });

const { onSearch, querier } = defineProps<{
  onSearch: (query: QueryPagi<any, any>) => void;
  querier: ReturnType<typeof useQueryForm>;
}>();

const form = computed(() => querier.form);

const OperationLabel: Record<Operation, MaybeRef<string>> = {
  [Operation.Contains]: computed(() => $t('operator.contains')),
  [Operation.ContainsFold]: computed(() => $t('operator.containsFold')),
  [Operation.Equal]: computed(() => $t('operator.equal')),
  [Operation.Gt]: computed(() => $t('operator.gt')),
  [Operation.Gte]: computed(() => $t('operator.gte')),
  [Operation.HasPrefix]: computed(() => $t('operator.hasPrefix')),
  [Operation.HasSuffix]: computed(() => $t('operator.hasSuffix')),
  [Operation.In]: computed(() => $t('operator.in')),
  [Operation.IsNull]: computed(() => $t('operator.isNull')),
  [Operation.Lt]: computed(() => $t('operator.lt')),
  [Operation.Lte]: computed(() => $t('operator.lte')),
  [Operation.NotEqual]: computed(() => $t('operator.notEqual')),
  [Operation.NotIn]: computed(() => $t('operator.notIn')),
  [Operation.Between]: computed(() => $t('operator.between')),
  [Operation.BetweenEq]: computed(() => $t('operator.betweenEq')),
};

const filterColumnOptions = computed(() => {
  return querier.form.columns.filter((c) => querier.types.get(c));
});

const existedFilters = computed(() => {
  const exists: Record<string, boolean> = {};
  form.value.filters.forEach((f) => {
    exists[String(f.column)] = true;
  });

  return exists;
});

const shouldDisplayButton = computed(() => {
  const filter =
    querier.config.button?.isFilter && filterColumnOptions.value.length > 0;

  const searchInput = querier.config.button?.isSearch;
  const reset = querier.config.button?.isReset;
  const sort = querier.sortOption.length > 0;

  return {
    filter,
    reset,
    searchInput,
    sort,
  };
});

const handleChangeColumn = (column: string, index: number) => {
  const filter = querier.form.filters?.[index];

  if (!filter) {
    return;
  }

  filter.column = column;
  filter.type = querier.types.get(column) || FieldType.string;

  switch (filter.type) {
    case FieldType.date: {
      filter.value = undefined;
      break;
    }

    case FieldType.select: {
      filter.operation = Operation.In;
      filter.value = [];
      break;
    }

    default: {
      filter.operation = Operation.Equal;
      filter.value = '';
      break;
    }
  }
};

const handleSortChange = (v: string[]) => {
  querier.setForm((f) => {
    const kv = new Map<string, string>();
    v.forEach((str) => {
      const [k, v] = str.split(' ');
      if (k && v) {
        kv.set(k, v);
      }
    });

    f.sorts = kv
      .entries()
      .map(([k, v]) => `${k} ${v}`)
      .toArray();
    return f;
  });
};

const getFieldConfig = ref(querier.getFieldConfig);
const filterNumber = computed(() => querier.form.filters.length);
const filterPopoverVisible = ref<boolean>(false);
const hideFilterPopover = () => {
  filterPopoverVisible.value = false;
};

const handleRemove = (index: number) => {
  querier.setForm((f) => {
    f.filters = f.filters.filter((_: any, i: number) => i !== index) ?? [];
    return f;
  });
};

const filterInputId = (column: any) => `filter-input-id-${column}`;

const handleAdd = (column: string) => {
  hideFilterPopover();

  const baseFilter = querier.schema.filters.find((e) => e.column === column);
  if (!baseFilter) return;

  querier.setForm((form) => {
    const newFilter = { ...baseFilter, column };
    form.filters.push(newFilter);

    setTimeout(() => {
      const inputEl = document.querySelector<HTMLInputElement>(
        `#${filterInputId(column)}`,
      );
      inputEl?.focus?.();
    }, 300);

    return form;
  });
};

const handleSubmit = () => {
  querier.setQuery((q) => {
    q.page = 1;
    return q;
  });

  onSearch(querier.getQuery());
};

const handleReset = () => {
  querier.resetForm();
  handleSubmit();
};
const rangePresets = computed<PresetDate<Dayjs[]>[]>(() => [
  {
    label: $t('common.last5Minutes'),
    value: [dateUtil().add(-5, 'minute'), dateUtil()],
  },
  {
    label: $t('common.last30Minutes'),
    value: [dateUtil().add(-30, 'minute'), dateUtil()],
  },
  {
    label: $t('common.today'),
    value: [dateUtil().startOf('d'), dateUtil()],
  },
  {
    label: $t('common.yesterday'),
    value: [
      dateUtil().add(-1, 'd').startOf('d'),
      dateUtil().add(-1, 'd').endOf('d'),
    ],
  },
  {
    label: $t('common.last7Days'),
    value: [dateUtil().add(-7, 'd').startOf('d'), dateUtil()],
  },
  {
    label: $t('common.last14Days'),
    value: [dateUtil().add(-14, 'd').startOf('d'), dateUtil()],
  },
  {
    label: $t('common.last30Days'),
    value: [dateUtil().add(-30, 'd').startOf('d'), dateUtil()],
  },
  {
    label: $t('common.last90Days'),
    value: [dateUtil().add(-90, 'd').startOf('d'), dateUtil()],
  },
]);
</script>

<template>
  <div class="filter-form">
    <Row>
      <Col class="space-y-2" flex="4">
        <Row :gutter="6" class="gap-y-2">
          <Col v-if="shouldDisplayButton.searchInput" style="width: 300px">
            <Input
              v-model:value="form.search.value"
              :on-press-enter="handleSubmit"
              :placeholder="$t('common.search')"
              type="text"
            >
              <template #addonBefore>
                <Select
                  v-model:value="form.search.operation"
                  default-value="containsFold"
                  style="width: fit-content"
                >
                  <SelectOption
                    v-for="[key, value] in Object.entries(StringOp)"
                    :key="value"
                    :value="value"
                  >
                    {{ OperationLabel[value] || key }}
                  </SelectOption>
                </Select>
              </template>
            </Input>
          </Col>
          <Col>
            <Button style="width: 100px" type="primary" @click="handleSubmit()">
              <template #icon><Search class="anticon size-4" /></template>
              {{ $t('common.search') }}
            </Button>
          </Col>
          <Col v-if="shouldDisplayButton.filter">
            <Popover
              v-model:open="filterPopoverVisible"
              placement="bottomRight"
              trigger="click"
            >
              <Button class="flex items-center">
                <Add v-if="filterNumber <= 0" class="anticon size-4" />
                <Badge v-else :count="filterNumber" class="pr-2" />
                {{ $t('common.filter') }}
              </Button>
              <template #content>
                <Flex gap="small" vertical>
                  <Button
                    v-for="column in filterColumnOptions"
                    :key="column"
                    :disabled="existedFilters[String(column)]"
                    @click="handleAdd(String(column))"
                  >
                    {{ getFieldConfig(String(column))?.label }}
                  </Button>
                </Flex>
              </template>
            </Popover>
          </Col>
          <Col v-if="shouldDisplayButton.sort" class="filter-col">
            <Select
              :placeholder="$t('common.sort')"
              :value="form.sorts"
              class="min-w-40"
              mode="multiple"
              @change="(v) => handleSortChange(v as string[])"
            >
              <SelectOption
                v-for="{ label, value, direction } in querier.sortOption"
                :key="value"
                :value="value"
              >
                <div class="flex items-center space-x-1">
                  <span>{{ label }}</span>
                  <VbenIcon
                    v-if="direction === 'desc'"
                    icon="ant-design:caret-down-outlined"
                  />
                  <VbenIcon v-else icon="ant-design:caret-up-outlined" />
                </div>
              </SelectOption>
            </Select>
          </Col>

          <Col v-if="shouldDisplayButton.reset" class="filter-col" span="2.5">
            <Button
              class="flex items-center justify-center gap-2"
              danger
              type="primary"
              @click="handleReset"
            >
              <template #icon>
                <VbenIcon icon="lucide:x" />
              </template>
              {{ $t('common.reset') }}
            </Button>
          </Col>
        </Row>
        <ul class="filter-list flex flex-col gap-2" v-auto-animate>
          <li
            v-for="(item, index) in form.filters"
            :key="`${String(item?.column)}${index}`"
            class="filter-item"
          >
            <Row :gutter="6">
              <Col style="width: 140px">
                <Select
                  v-model:value="item.column as string"
                  :dropdown-match-select-width="false"
                  class="w-full"
                  @change="(c) => handleChangeColumn(c as string, index)"
                >
                  <SelectOption
                    v-for="column in filterColumnOptions"
                    :key="column"
                    :value="String(column)"
                  >
                    {{ getFieldConfig(String(column))?.label }}
                  </SelectOption>
                </Select>
              </Col>
              <template v-if="item.type === FieldType.number">
                <Col style="width: 163px">
                  <Select
                    v-model:value="item.operation"
                    :dropdown-match-select-width="false"
                    class="w-full"
                  >
                    <SelectOption
                      v-for="[key, value] in Object.entries(NumberOp)"
                      :key="value"
                      :value="value"
                    >
                      {{ OperationLabel[value] || key }}
                    </SelectOption>
                  </Select>
                </Col>
                <Col style="width: 163px">
                  <Input
                    :id="filterInputId(item.column)"
                    v-model:value="item.value"
                    :on-press-enter="handleSubmit"
                    placeholder="100"
                    type="number"
                  />
                </Col>
              </template>
              <template v-if="item.type === FieldType.string">
                <Col style="width: 163px">
                  <Select v-model:value="item.operation" class="w-full">
                    <SelectOption
                      v-for="[key, value] in Object.entries(StringOp)"
                      :key="value"
                      :value="value"
                    >
                      {{ OperationLabel[value] || key }}
                    </SelectOption>
                  </Select>
                </Col>
                <Col style="width: 163px">
                  <Input
                    :id="filterInputId(item.column)"
                    v-model:value="item.value"
                    :on-press-enter="handleSubmit"
                    :placeholder="$t('common.text')"
                    type="text"
                  />
                </Col>
              </template>
              <template v-if="item.type === FieldType.date">
                <Col>
                  <RangePicker
                    :id="filterInputId(item.column)"
                    v-model:value="item.value as any"
                    :disabled-date="disableFutureDates"
                    :presets="rangePresets"
                    format="YYYY-MM-DD HH:mm:ss"
                    show-time
                    style="width: 368px"
                  />
                </Col>
              </template>
              <template v-if="item.type === FieldType.select">
                <Col style="width: 163px">
                  <Select v-model:value="item.operation" class="w-full">
                    <SelectOption
                      v-for="[key, value] in Object.entries(SelectOp)"
                      :key="value"
                      :value="value"
                    >
                      {{ OperationLabel[value] || key }}
                    </SelectOption>
                  </Select>
                </Col>
                <Col style="min-width: 163px">
                  <Select
                    :id="filterInputId(item.column)"
                    v-model:value="item.value"
                    :dropdown-match-select-width="false"
                    :options="
                      getGetter(querier.selectOption.get(item.column))?.map(
                        (item) => ({
                          ...item,
                          value: item.value.toString(),
                        }),
                      )
                    "
                    allow-clear
                    class="w-full"
                    mode="multiple"
                    option-filter-prop="label"
                  />
                </Col>
              </template>
              <template v-if="!querier.getFieldConfig(item.column)?.isFixed">
                <Col>
                  <Button
                    class="flex items-center justify-center"
                    type="ghost"
                    @click="() => handleRemove(index)"
                  >
                    <template #icon>
                      <VbenIcon
                        class="text-destructive text-lg"
                        icon="lucide:trash-2"
                      />
                    </template>
                  </Button>
                </Col>
              </template>
            </Row>
          </li>
        </ul>
      </Col>
      <Col align="right" flex="1"> <slot></slot> </Col>
    </Row>
  </div>
</template>

<style lang="css">
.filter-form {
  display: flex;
  flex-direction: column;
  width: 100%;
  margin-bottom: 8px;
}

.filter-list {
  list-style-type: none;
}

.filter-item {
  list-style-type: none;
}
</style>
