# boringproxy.io

This repository contains the source files for the boringproxy.io website.

Build it with

```bash
npm i
npm run build          # builds into the current directory
npm run build docs     # builds into the docs subdirectory
```

If your system comes with the `entr` command, you can rebuild the site on file changes.

```bash
npm run watch          # builds and watches for file changes
npm run watch docs     # builds to docs and watches for file changes
```

Sometimes you need to start your build from scratch.

```bash
npm run clean          # cleans selected files in the current directory
npm run clean docs     # cleans the docs directory
```

See all build jobs and their specifications with `npm run.`

## License

This work is released under the ISC license.
