import Lenis from '@studio-freight/lenis'

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.hook('app:mounted', () => {
    const lenis = new Lenis()

    nuxtApp.vueApp.config.globalProperties.$lenis = lenis

    function raf(time: number) {
      lenis.raf(time)
      requestAnimationFrame(raf)
    }

    requestAnimationFrame(raf)
  })
})
