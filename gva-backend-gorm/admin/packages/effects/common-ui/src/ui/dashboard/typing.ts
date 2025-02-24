import type { Component } from 'vue';

enum AnalysisOverviewItemKey {
  REQUEST_SUCCESS_RATE = 'requestSuccessRate',
  TOTAL_REQUEST = 'totalRequest',
  TOTAL_USER_GAIN = 'totalUserGain',
  TOTAL_VISIT = 'totalVisit',
}

interface AnalysisOverviewItem {
  icon: Component | string;
  title: string;
  suffix?: string;
  prefix?: string;
  key: AnalysisOverviewItemKey;
  color?: string;
}

interface WorkbenchProjectItem {
  color?: string;
  content: string;
  date: string;
  group: string;
  icon: Component | string;
  title: string;
}

interface WorkbenchTrendItem {
  avatar: string;
  content: string;
  date: string;
  title: string;
}

interface WorkbenchTodoItem {
  completed: boolean;
  content: string;
  date: string;
  title: string;
}

interface WorkbenchQuickNavItem {
  color?: string;
  icon: Component | string;
  title: string;
}

export type {
  AnalysisOverviewItem,
  WorkbenchProjectItem,
  WorkbenchQuickNavItem,
  WorkbenchTodoItem,
  WorkbenchTrendItem,
};

export { AnalysisOverviewItemKey };
