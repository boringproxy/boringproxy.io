#!/usr/bin/env node

const fs = require('fs');
const path = require('path');
const marked = require('marked');
const hljs = require('highlightjs');

marked.setOptions({
  highlight: function(code, lang) {
    const highlighted = hljs.highlightAuto(code);
    return highlighted.value;
  },
});

const inDir = './';
const outDir = (process.argv.length !== 3) ? './' : `./${process.argv[2]}/`;
fs.mkdirSync(outDir, { recursive: true });

const partialsDir = path.join(inDir, 'partials');

const headerHtml = fs.readFileSync(path.join(partialsDir, 'header.html'), 'utf-8');
const footerHtml = fs.readFileSync(path.join(partialsDir, 'footer.html'), 'utf-8');

const pagesDir = path.join(inDir, 'pages');
const pagesFiles = fs.readdirSync(pagesDir, {
  withFileTypes: true,
});

for (const pageMd of pagesFiles) {
  if (!pageMd.isFile()) {
    continue;
  }

  const mdText = fs.readFileSync(path.join(pagesDir, pageMd.name), 'utf-8');
  const htmlText = headerHtml + marked(mdText) + footerHtml;
  const pageName = path.parse(pageMd.name).name;

  let htmlDir;
  if (pageName === 'index') {
    htmlDir = outDir;
  }
  else {
    htmlDir = path.join(outDir, pageName);
    fs.mkdirSync(htmlDir, { recursive: true });
  }

  fs.writeFileSync(path.join(htmlDir, 'index.html'), htmlText);
}

if (outDir != './') {
  for (file of ["styles.css", "logo.svg", "screenshot.png"]) {
    fs.copyFileSync(`${file}`, path.join(outDir, file))
  }
}
