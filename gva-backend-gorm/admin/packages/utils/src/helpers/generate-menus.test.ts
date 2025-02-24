import {
  createRouter,
  createWebHistory,
  type Router,
  type RouteRecordRaw,
} from 'vue-router';

import { describe, expect, it, vi } from 'vitest';

import { generateMenus } from './generate-menus'; // 替换为您的实际路径

// Nested route setup to test child inclusion and hideChildrenInMenu functionality

describe('generateMenus', () => {
  // 模拟路由数据
  const mockRoutes = [
    {
      meta: { icon: 'home-icon', title: '首页' },
      name: 'home',
      path: '/home',
    },
    {
      children: [
        {
          meta: { icon: 'team-icon', title: '团队' },
          name: 'team',
          path: 'team',
        },
      ],
      meta: { hideChildrenInMenu: true, icon: 'about-icon', title: '关于' },
      name: 'about',
      path: '/about',
    },
  ] as RouteRecordRaw[];

  // 模拟 Vue 路由器实例
  const mockRouter = {
    getRoutes: vi.fn(() => [
      { name: 'home', path: '/home' },
      { name: 'about', path: '/about' },
      { name: 'team', path: '/about/team' },
    ]),
  };

  it('the correct menu list should be generated according to the route', async () => {
    const expectedMenus = [
      {
        badge: undefined,
        badgeType: undefined,
        badgeVariants: undefined,
        children: [],
        icon: 'home-icon',
        name: '首页',
        order: undefined,
        parent: undefined,
        parents: undefined,
        path: '/home',
        show: true,
      },
      {
        badge: undefined,
        badgeType: undefined,
        badgeVariants: undefined,
        children: [],
        icon: 'about-icon',
        name: '关于',
        order: undefined,
        parent: undefined,
        parents: undefined,
        path: '/about',
        show: true,
      },
    ];

    const menus = await generateMenus(mockRoutes, mockRouter as any);
    expect(menus).toEqual(expectedMenus);
  });

  it('includes additional meta properties in menu items', async () => {
    const mockRoutesWithMeta = [
      {
        meta: { icon: 'user-icon', order: 1, title: 'Profile' },
        name: 'profile',
        path: '/profile',
      },
    ] as RouteRecordRaw[];

    const menus = await generateMenus(mockRoutesWithMeta, mockRouter as any);
    expect(menus).toEqual([
      {
        badge: undefined,
        badgeType: undefined,
        badgeVariants: undefined,
        children: [],
        icon: 'user-icon',
        name: 'Profile',
        order: 1,
        parent: undefined,
        parents: undefined,
        path: '/profile',
        show: true,
      },
    ]);
  });

  it('handles dynamic route parameters correctly', async () => {
    const mockRoutesWithParams = [
      {
        meta: { icon: 'details-icon', title: 'User Details' },
        name: 'userDetails',
        path: '/users/:userId',
      },
    ] as RouteRecordRaw[];

    const menus = await generateMenus(mockRoutesWithParams, mockRouter as any);
    expect(menus).toEqual([
      {
        badge: undefined,
        badgeType: undefined,
        badgeVariants: undefined,
        children: [],
        icon: 'details-icon',
        name: 'User Details',
        order: undefined,
        parent: undefined,
        parents: undefined,
        path: '/users/:userId',
        show: true,
      },
    ]);
  });

  it('processes routes with redirects correctly', async () => {
    const mockRoutesWithRedirect = [
      {
        name: 'redirectedRoute',
        path: '/old-path',
        redirect: '/new-path',
      },
      {
        meta: { icon: 'path-icon', title: 'New Path' },
        name: 'newPath',
        path: '/new-path',
      },
    ] as RouteRecordRaw[];

    const menus = await generateMenus(
      mockRoutesWithRedirect,
      mockRouter as any,
    );
    expect(menus).toEqual([
      // Assuming your generateMenus function excludes redirect routes from the menu
      {
        badge: undefined,
        badgeType: undefined,
        badgeVariants: undefined,
        children: [],
        icon: undefined,
        name: 'redirectedRoute',
        order: undefined,
        parent: undefined,
        parents: undefined,
        path: '/old-path',
        show: true,
      },
      {
        badge: undefined,
        badgeType: undefined,
        badgeVariants: undefined,
        children: [],
        icon: 'path-icon',
        name: 'New Path',
        order: undefined,
        parent: undefined,
        parents: undefined,
        path: '/new-path',
        show: true,
      },
    ]);
  });

  const routes: any = [
    {
      meta: { order: 2, title: 'Home' },
      name: 'home',
      path: '/',
    },
    {
      meta: { order: 1, title: 'About' },
      name: 'about',
      path: '/about',
    },
  ];

  const router: Router = createRouter({
    history: createWebHistory(),
    routes,
  });

  it('should generate menu list with correct order', async () => {
    const menus = await generateMenus(routes, router);
    const expectedMenus = [
      {
        badge: undefined,
        badgeType: undefined,
        badgeVariants: undefined,
        children: [],
        icon: undefined,
        name: 'About',
        order: 1,
        parent: undefined,
        parents: undefined,
        path: '/about',
        show: true,
      },
      {
        badge: undefined,
        badgeType: undefined,
        badgeVariants: undefined,
        children: [],
        icon: undefined,
        name: 'Home',
        order: 2,
        parent: undefined,
        parents: undefined,
        path: '/',
        show: true,
      },
    ];

    expect(menus).toEqual(expectedMenus);
  });

  it('should handle empty routes', async () => {
    const emptyRoutes: any[] = [];
    const menus = await generateMenus(emptyRoutes, router);
    expect(menus).toEqual([]);
  });
});
