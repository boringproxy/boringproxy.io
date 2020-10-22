#!/usr/bin/env node

const fs = require('fs');
const path = require('path');
const marked = require('marked');

const inDir = './';
const outDir = './';

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
  const htmlFilename = path.parse(pageMd.name).name + '.html';
  fs.writeFileSync(path.join(outDir, htmlFilename), htmlText);
}
