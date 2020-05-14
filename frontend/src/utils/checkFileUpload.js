export function streamFiles(ipfs, files) {
  return new Promise((resolve, reject) => {
    const stream = ipfs.addReadableStream({
      wrapWithDirectory: true
      // progress: (length: number) =>
      //     setFileSizeReceived(formatBytes(length, 0))
    });

    stream.on('data', (data) => {
      // The last data event will contain the directory hash
      if (data.path === '') resolve(data.hash);
    });

    stream.on('error', reject);
    stream.write(files);
    stream.end();
  });
}
