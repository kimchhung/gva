import dayjs, { type Dayjs } from 'dayjs';
import timezone from 'dayjs/plugin/timezone';
import utc from 'dayjs/plugin/utc';

dayjs.extend(utc);
dayjs.extend(timezone);
dayjs.tz.setDefault('Asia/Shanghai');

const dateUtil = dayjs;

const DATE_TIME_FORMAT = 'YYYY-MM-DD HH:mm:ss';
const DATE_TIME_WITHOUT_YEAR_FORMAT = 'MM-DD HH:mm:ss';
const DATE_FORMAT = 'YYYY-MM-DD';

function formatToDateTime(
  date?: dayjs.ConfigType,
  format = DATE_TIME_FORMAT,
): string {
  const now = dateUtil();
  const currentYear = now.year();
  const dateYear = dateUtil(date).year();

  if (currentYear === dateYear) {
    return dateUtil(date).format(DATE_TIME_WITHOUT_YEAR_FORMAT);
  }

  return dateUtil(date).format(format);
}

function formatDateToString(date?: dayjs.ConfigType): string {
  return dateUtil(date).format();
}

function formatToDate(date?: dayjs.ConfigType, format = DATE_FORMAT): string {
  return dateUtil(date).format(format);
}

const getTimeHasPassed = (
  date: Dayjs,
  now: Date | string,
  format?: string,
): string => {
  const r = Math.abs(date.diff(now, 'minute'));

  if (r < 5) {
    return '现在';
  }
  if (r >= 60 && r < 60 * 24) {
    return `${(r / 60).toFixed(0)}小时前`;
  }
  if (r >= 60 * 24 && r < 60 * 24 * 3) {
    return `${(r / (60 * 24)).toFixed(0)}天前`;
  }
  if (r >= 60 * 24 * 3) {
    return formatToDateTime(date, format ?? 'MM月DD日');
  }

  return `${r}分钟前`;
};

export const todayOnly = (): [Dayjs, Dayjs] => [
  dateUtil().startOf('day'),
  dateUtil().endOf('day'),
];

export const yesterdayOnly = (): [Dayjs, Dayjs] => [
  dateUtil().subtract(1, 'day').startOf('day'),
  dateUtil().subtract(1, 'day').endOf('day'),
];

export const last2days = (): [Dayjs, Dayjs] => [
  dateUtil().subtract(2, 'day').startOf('day'),
  dateUtil().subtract(2, 'day').endOf('day'),
];

const disableFutureDates = (current: Dayjs) => {
  return current && current.isAfter(dateUtil().endOf('day'));
};

export {
  DATE_FORMAT,
  DATE_TIME_FORMAT,
  dateUtil,
  disableFutureDates,
  formatDateToString,
  formatToDate,
  formatToDateTime,
  getTimeHasPassed,
};
