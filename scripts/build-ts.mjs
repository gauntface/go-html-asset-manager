import * as esbuild from 'esbuild';

const entries = [
  {
    entry: 'static/js/components/n-ham-c-lite-yt-async.ts',
    outfile: 'embedassets/assets/js/components/n-ham-c-lite-yt-async.js',
  },
  {
    entry: 'static/js/components/n-ham-c-lite-vi-async.ts',
    outfile: 'embedassets/assets/js/components/n-ham-c-lite-vi-async.js',
  },
  {
    entry: 'static/js/bootstrap/always-async.ts',
    outfile: 'embedassets/assets/js/bootstrap/always-async.js',
  },
];

async function main() {
  for (const {entry, outfile} of entries) {
    await esbuild.build({
      entryPoints: [entry],
      bundle: true,
      minify: true,
      format: 'cjs',
      outfile,
    });
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
