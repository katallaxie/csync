// theme.config.js
export default {
    github: 'https://github.com/katallaxie/csync',
    docsRepositoryBase: 'https://github.com/katallaxie/csync/blob/main/docs/pages', // base URL for the docs repository
    titleSuffix: ' – csync',
    nextLinks: true,
    prevLinks: true,
    search: true,
    customSearch: null, // customizable, you can use algolia for example
    darkMode: true,
    footer: true,
    footerText: `MIT ${new Date().getFullYear()} © Sebastian Doell (@katallaxie).`,
    footerEditLink: `Edit this page on GitHub`,
    logo: (
      <>
        <svg>...</svg>
        <span>csync</span>
      </>
    )
}