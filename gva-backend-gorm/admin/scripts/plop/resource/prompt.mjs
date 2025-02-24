import { exec } from 'child_process';

export default function definePlop(
  /** @type {import('plop').NodePlopAPI} */
  plop,
) {
  plop.setGenerator('resource', {
    description: 'Generate a resource',
    prompts: [
      {
        name: 'name',
        type: 'input',
      },
      {
        name: 'viewPath',
        type: 'input',
        default: '/views',
      },
    ],

    actions: [
      {
        // folder copy base type=addMany
        // the type of action (add, modify, addMany)
        type: 'addMany',
        destination: 'app/src/api/{{dashCase name}}',
        templateFiles: 'scripts/plop/resource/templates/api/*.tmpl',
        base: 'scripts/plop/resource/templates/api',
        stripExtensions: ['tmpl'],
      },

      {
        // folder name single avoid mess up between plural (y-ies)
        type: 'addMany',
        destination: 'app/src/views/{{viewPath}}/{{dashCase name}}',
        templateFiles: 'scripts/plop/resource/templates/views/*.tmpl',
        base: 'scripts/plop/resource/templates/views',
        stripExtensions: ['tmpl'],
        data: '{}',
      },
      function () {
        console.log('pulling permission');
        exec('pnpm pull:permission', (error, stdout, stderr) => {
          if (error) {
            console.error(`exec error: ${error}`);
            throw new Error(error);
          }
          if (stderr) {
            console.error(`stderr: ${stderr}`);
          }
          console.log(stdout);
        });
        return 'pnpm command executed';
      },
    ],
  });
}
