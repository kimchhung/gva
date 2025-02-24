import type { CAC } from 'cac';
import fs from 'node:fs';
async function pullPermission({ url, path }: { url: string; path: string }) {
  const response = await fetch(url);
  const resp = (await response.json()) as { data: string[] };

  /*
  export enum ADMIN_ROLE_PERMISSION {
    ADD = 'admin_role:add',
    DELETE = 'admin_role:delete',
    EDIT = 'admin_role:edit',
    SUPER = 'admin_role:super',
    VIEW = 'admin_role:view',
  }
   */
  const enums: Record<string, Record<string, string>> = {};
  resp.data?.forEach((scope) => {
    const [group, action] = scope.split(':');
    if (!group || !action) return;

    const enumName = group.toUpperCase() + '_PERMISSION';
    if (!enums[enumName]) enums[enumName] = {} as Record<string, string>;
    enums[enumName][action.toUpperCase()] = `'${scope}'`;
  });

  for (const enumName in enums) {
    if (!enums[enumName]) continue;
    const enumValues = Object.entries(enums[enumName]);
    enumValues.sort((a, b) => a[0].localeCompare(b[0]));
    enums[enumName] = Object.fromEntries(enumValues);
  }

  const enumLines = Object.entries(enums)
    .sort((a, b) => a[0].localeCompare(b[0]))
    .map(([key, value]) => {
      return `export enum ${key} {\n${Object.entries(value)
        .map(([key, value]) => `  ${key} = ${value},`)
        .join('\n')}\n}`;
    });

  const startLine = `/* ------------------------------
  ! Generated files, do not edit
  ! Use pnpm run pull:permission instead
  ------------------------------ */\n\n`;

  let line = startLine + enumLines.join('\n\n');
  line += '\n\n';
  line += 'export type PERMISSION =\n';
  line += '  | ';
  line += `${Object.keys(enums)
    .sort((a, b) => a.localeCompare(b))
    .join('\n  | ')};\n`;

  fs.writeFile(
    path,
    line,
    { encoding: 'utf-8', flag: 'w', mode: 0o666 },
    (err) => {
      if (err) console.error(err);
    },
  );
}

function definePullPermissionCommand(cac: CAC) {
  cac
    .command('pull-permission')
    .usage('pull permission and write to file')
    .option('--url <url>', 'add url fetch permission')
    .option('--path <path>', 'path write file into')
    .action(async ({ url, path }) => {
      await pullPermission({ url, path });
    });
}

export { definePullPermissionCommand };
