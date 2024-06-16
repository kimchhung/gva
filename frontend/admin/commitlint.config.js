module.exports = {
  extends: ['@commitlint/config-conventional'],
  rules: {
    'type-enum': [
      2,
      'always',
      [
        'feat', // (feature)
        'fix', // (bug)
        'docs', // (documentation)
        'style', // (Formatting, style)
        'refactor', // (Code refactoring)
        'perf', // (Performance optimization)
        'test', // (Adding test)
        'ci', // (Continuous Integration changes)
        'chore', // (Build process or auxiliary tool changes)
        'revert', // (Reverting to the previous version)
        'workflow', // (Workflow improvements)
        'mod', // (Uncertain category modifications)
        'wip', // (Work in progress)
        'types', // (Type modifications)
        'release' // (Version release)
      ]
    ],
    'subject-full-stop': [0, 'never'],
    'subject-case': [0, 'never']
  }
}
