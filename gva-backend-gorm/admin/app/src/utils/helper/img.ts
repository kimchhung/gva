export const displayLegacyImage = (url: string) => {
  const tempUrl = 'https://one-file-upload.my-captain.ss.hd1.fun/files';
  if (!url.includes('http')) {
    return `${tempUrl}${url}`;
  }
  return url;
};
