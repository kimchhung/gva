import chalk from 'chalk'
import fs from 'fs-extra'
import inquirer from 'inquirer'
import path from 'path'
import pkg from '../package.json'

async function generateIcon() {
  const dir = path.resolve(process.cwd(), 'node_modules/@iconify/json')

  const raw = await fs.readJSON(path.join(dir, 'collections.json'))

  const collections = Object.entries(raw).map(([id, v]) => ({
    ...(v as any),
    id
  }))

  const choices = collections.map((item) => ({ key: item.id, value: item.id, name: item.name }))

  inquirer
    .prompt([
      // {
      //   type: 'list',
      //   name: 'useType',
      //   choices: [
      //     { key: 'local', value: 'local', name: 'Local' },
      //     { key: 'onLine', value: 'onLine', name: 'OnLine' }
      //   ],
      //   message: 'How to use icons?'
      // },
      {
        type: 'list',
        name: 'iconSet',
        choices: choices,
        message: 'Select the icon set that needs to be generated?'
      }
    ])
    // ↓命令行问答的答案
    .then(async (answers) => {
      const { iconSet } = answers
      // const isOnLine = useType === 'onLine'
      const outputDir = path.resolve(process.cwd(), 'src/components/IconPicker/src/data')
      fs.ensureDir(outputDir)
      const genCollections = collections.filter((item) => [iconSet].includes(item.id))
      const prefixSet: string[] = []
      for (const info of genCollections) {
        const data = await fs.readJSON(path.join(dir, 'json', `${info.id}.json`))
        if (data) {
          const { prefix } = data
          const icons = Object.keys(data.icons).map((item) => `${prefix}:${item}`)

          await fs.writeFileSync(
            path.join('src/components/IconPicker/src/data', `icons.${prefix}.ts`),
            `export default ${JSON.stringify({ name: info.name, prefix, icons })}`
          )
          // ↓ Classified processing is complete, push type name
          prefixSet.push(prefix)
        }
      }
      console.log(
        `✨ ${chalk.cyan(`[${pkg.name}]`)}` + ' - Icon generated successfully:' + `[${prefixSet}]`
      )
    })
}

generateIcon()
