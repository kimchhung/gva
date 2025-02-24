/**
 * Global authority directive
 * Used for fine-grained control of component permissions
 * @Example v-auth="RoleEnum.TEST"
 */
import type { App, Directive, DirectiveBinding } from 'vue';

import { useAccess } from './use-access';

function isPermissionAccessible(el: Element, binding: any) {
  const { hasAccessByPermissions } = useAccess();

  const value = binding.value;

  if (!value) {
    return;
  }

  if (!hasAccessByPermissions(value)) {
    el?.remove();
  }
}

function isRoleAccessible(el: Element, binding: any) {
  const { hasAccessByRoles } = useAccess();

  const value = binding.value;

  if (!value) {
    return;
  }

  if (!hasAccessByRoles(value)) {
    el?.remove();
  }
}

const permissionDirective: Directive = {
  mounted: (el: Element, binding: DirectiveBinding<string | string[]>) => {
    isPermissionAccessible(el, binding);
  },
};

const roleDirective: Directive = {
  mounted: (el: Element, binding: DirectiveBinding<string | string[]>) => {
    isRoleAccessible(el, binding);
  },
};

export function registerAccessDirective(app: App) {
  app.directive('permissions', permissionDirective);
  app.directive('roles', roleDirective);
}
