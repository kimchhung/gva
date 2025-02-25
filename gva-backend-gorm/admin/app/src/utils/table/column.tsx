import type { DefaultOptionType } from 'ant-design-vue/es/select';

import { type MaybeRef, unref } from 'vue';

import { $t } from '@vben/locales';
import { VbenIcon } from '@vben-core/shadcn-ui';

import {
  Image,
  Space,
  type TableColumnType,
  Tag,
  type TagProps,
} from 'ant-design-vue';

import ConfirmationSelect from '#/views/_core/base/confirmation-select.vue';

import { copyToClipboard } from '../helper/copy';
import { formatToDateTime } from '../helper/date-util';
import { displayLegacyImage } from '../helper/img';

type ColumnConfigFunction = (config?: TableColumnType) => TableColumnType;

type OptionTypeWithColor = {
  color?: string;
} & DefaultOptionType;

const nullText = '-/-';

const getKey = (dataIndex: string | string[]): string => {
  return Array.isArray(dataIndex) ? dataIndex.join('-') : dataIndex;
};

const textColumn = (
  dataIndex: string | string[],
  title: string,
  other?: TableColumnType,
): TableColumnType => {
  return {
    dataIndex,
    key: getKey(dataIndex),
    title,
    customRender: ({ text }) => {
      return text || nullText;
    },
    ...other,
  };
};

const dateColumn = (
  dataIndex: string,
  title: string,
  other?: TableColumnType,
): TableColumnType => {
  return {
    dataIndex,
    title,
    customRender: ({ text }: { text: string }) => {
      if (!text) return nullText;

      return formatToDateTime(text);
    },
    width: 180,
    ...other,
  };
};

const tagColumn = (
  dataIndex: string | string[],
  title: string,
  other?: {
    getItem?: (item: any) => null | number | string | undefined;
    getProps?: (item: any) => TagProps;
  } & TableColumnType,
): TableColumnType => {
  return {
    dataIndex,
    title,
    key: getKey(dataIndex),
    customRender: ({ value }) => {
      let _value = value;
      if (
        typeof value === 'string' ||
        typeof value === 'number' ||
        typeof value === 'boolean'
      ) {
        _value = [value];
      }

      if (!Array.isArray(_value)) return nullText;
      if (_value.length === 0) return nullText;
      if (typeof value === 'string' && value === '') return nullText;

      return (
        <Space style={{ flexWrap: 'wrap' }}>
          {_value.map((item: string) => {
            let content = null;
            content = other?.getItem ? other.getItem(item) : item;

            if (content === null || content === undefined) return nullText;
            return (
              <Tag
                {...(other?.getProps ? other.getProps(item) : {})}
                key={content}
              >
                {content}
              </Tag>
            );
          })}
        </Space>
      );
    },
    ...other,
  };
};

const action: ColumnConfigFunction = (config) => {
  return {
    dataIndex: 'action',
    key: 'action',
    align: 'center',
    customHeaderCell: () => ({
      class: 'text-center',
    }),
    fixed: 'right',
    width: 200,
    title: $t('common.action'),
    ...config,
  };
};

const id: ColumnConfigFunction = (config) => {
  return textColumn('id', $t('common.id'), {
    width: 100,
    ...config,
  });
};

const index: ColumnConfigFunction = (config) => {
  return textColumn('index', '#', {
    width: 100,
    ...config,
  });
};

const name: ColumnConfigFunction = (config) => {
  return textColumn('name', $t('common.name'), config);
};

const timezone = (
  dataIndex: string | string[],
  other?: TableColumnType,
): TableColumnType => {
  return {
    dataIndex,
    key: getKey(dataIndex),
    title: $t('common.timezone'),
    customRender: ({ text }) => {
      const renderedText = text ? $t(`TIMEZONE.${text}`) : nullText;
      return renderedText.includes('TIMEZONE.') ? text : renderedText;
    },
    ...other,
  };
};

const createdAt: ColumnConfigFunction = (config) => {
  return dateColumn('createdAt', $t('common.createdAt'), config);
};

const updatedAt: ColumnConfigFunction = (config) => {
  return dateColumn('updatedAt', $t('common.updatedAt'), config);
};

const status = <TRecord = any,>(
  config?: {
    disabled?: ((record: TRecord) => boolean) | boolean;
    onConfirm?: (newValue: any, record: TRecord) => Promise<any>;
    options?: MaybeRef<OptionTypeWithColor[]>;
  } & TableColumnType<TRecord>,
): TableColumnType<TRecord> => {
  if (!config || !config.onConfirm) {
    return tagColumn('status', $t('common.status'), {
      getItem: (item) => {
        return item ? $t('common.enable') : $t('common.disable');
      },
      getProps: (item) => {
        return {
          color: item ? 'green' : 'red',
        };
      },
      ...config,
    });
  }

  return {
    dataIndex: 'status',
    key: 'status',
    width: 180,
    title: $t('common.status'),
    customRender: ({ value, record }) => {
      const options = config?.options ?? [
        { label: $t('common.enable'), value: 1, color: 'green' },
        { label: $t('common.disable'), value: 0, color: 'red' },
      ];

      const key =
        record && typeof record === 'object' && 'id' in record
          ? (record.id as string)
          : JSON.stringify(record);

      const isDisabled = !!(config?.disabled &&
      typeof config.disabled === 'function'
        ? config.disabled(record)
        : config.disabled);

      return (
        <ConfirmationSelect
          defaultValue={value}
          disabled={isDisabled}
          key={key}
          onConfirm={(value) =>
            config?.onConfirm && config?.onConfirm?.(value, record)
          }
          options={unref(options)}
        />
      );
    },
    ...config,
  };
};

function image<T>(
  config: {
    className?: string;
    dataIndex: string | string[];
    getSrc?: (record: T) => string;
    height?: number;
    title: string;
    width?: number;
  } & TableColumnType,
): TableColumnType {
  const { className, ...otherConfig } = config;
  return {
    customRender: ({ value, record }) => {
      let src = config.getSrc ? config.getSrc(record) : value;

      // TODO: remove after replace all image
      if (src) {
        src = displayLegacyImage(src);
      }

      const width = config.width ?? 50;
      const height = config.height ?? 50;

      if (!src) {
        return (
          <div
            class={'flex items-center justify-center rounded-sm border-[0.5px]'}
            style={{ width: `${width}px`, height: `${height}px` }}
          >
            <VbenIcon class="size-5" icon="lucide:image-off" />
          </div>
        );
      }

      return (
        <Image
          class={`rounded-sm ${className}`}
          height={height}
          src={src}
          width={width}
        />
      );
    },
    width: 120,
    ...otherConfig,
  };
}

const level = (config?: TableColumnType): TableColumnType => {
  return tagColumn('socialLevel', $t('common.socialLevel'), {
    getItem: (item) => {
      return `lvl.${item}`;
    },
    ...config,
  });
};

const copy = (config: TableColumnType): TableColumnType => {
  return {
    customRender: ({ text }) => {
      if (!text) return nullText;

      return (
        <div class="flex items-center space-x-2 break-all">
          <p class="line-clamp-2">{text}</p>
          <button onClick={() => copyToClipboard(text)}>
            <VbenIcon icon="carbon:copy" />
          </button>
        </div>
      );
    },
    ...config,
  };
};

const username = (config?: TableColumnType): TableColumnType => {
  return copy({
    dataIndex: 'username',
    title: $t('common.username'),
    ...config,
  });
};

const amount = (config?: TableColumnType): TableColumnType => {
  return {
    dataIndex: 'amount',
    key: 'amount',
    title: $t('common.amount'),
    customRender: ({ text }) => {
      const className = text > 0 ? 'text-green-500' : 'text-red-500';
      return (
        <p class={`${className} font-semibold`}>
          {text > 0 ? '+' : ''}
          {text}
        </p>
      );
    },
    ...config,
  };
};

const goldCoin = (
  config?: {
    getItem?: (item: any) => null | number | string | undefined;
  } & TableColumnType,
): TableColumnType => {
  return {
    dataIndex: 'goldCoin',
    key: 'goldCoin',
    title: $t('common.goldCoin'),
    width: 140,
    customRender: ({ text, value }) => {
      return (
        <p class="inline-flex items-center">
          <VbenIcon
            class="mr-1 size-5"
            icon="ant-design:pay-circle-outlined"
            style={{ color: '#FF8800' }}
          />
          {config?.getItem ? config.getItem(value) : text}
        </p>
      );
    },
    ...config,
  };
};

const expPoint = (
  config?: {
    getItem?: (item: any) => null | number | string | undefined;
  } & TableColumnType,
): TableColumnType => {
  return {
    dataIndex: 'expPoint',
    key: 'expPoint',
    title: $t('common.expPoint'),
    width: 140,
    customRender: ({ value, text }) => {
      return (
        <p class="inline-flex items-center">
          <VbenIcon
            class="mr-1 size-5"
            icon="ant-design:crown-outlined"
            style={{ color: '#00ccbb' }}
          />
          {config?.getItem ? config.getItem(value) : text}
        </p>
      );
    },
    ...config,
  };
};

export const removeColumn = (columns: any[], col: string) => {
  return columns.filter((column) => column.key !== col);
};

export const tableColumns = {
  action,
  id,
  name,
  status,
  createdAt,
  updatedAt,
  dateColumn,
  textColumn,
  tagColumn,
  image,
  index,
  level,
  copy,
  username,
  amount,
  nullText,
  goldCoin,
  expPoint,
  timezone,
};
