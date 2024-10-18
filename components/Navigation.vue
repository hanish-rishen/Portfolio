<template>
  <nav class="fixed top-0 left-0 w-full bg-black bg-opacity-80 backdrop-blur-sm z-50 py-4 px-6 border-b border-white border-opacity-20">
    <div class="container mx-auto flex justify-between items-center">
      <div class="text-lg font-medium" ref="timeDisplay"></div>
      <ul class="flex space-x-6">
        <li v-for="section in sections" :key="section.id">
          <a :href="`#${section.id}`" class="text-gray-300 hover:text-white transition-colors duration-300 relative group" @mouseenter="expandCursor" @mouseleave="shrinkCursor">
            {{ section.title }}
            <span class="absolute left-0 bottom-0 w-full h-0.5 bg-white transform scale-x-0 transition-transform duration-300 origin-left group-hover:scale-x-100"></span>
          </a>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { gsap } from 'gsap'

const sections = [
  { id: 'about', title: 'About Me', subtitle: 'A brief introduction about yourself' },
  { id: 'projects', title: 'Projects', subtitle: 'Showcase your best work' },
  { id: 'skills', title: 'Skills', subtitle: 'Your technical expertise' },
  { id: 'contact', title: 'Contact', subtitle: 'Get in touch' }
]

const timeDisplay = ref(null)
let timeInterval

const updateTime = () => {
  const now = new Date()
  timeDisplay.value.textContent = now.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
}

const expandCursor = () => {
  gsap.to('.custom-cursor', {
    scale: 1.5,
    duration: 0.3,
    ease: 'power2.out'
  })
}

const shrinkCursor = () => {
  gsap.to('.custom-cursor', {
    scale: 1,
    duration: 0.3,
    ease: 'power2.out'
  })
}

onMounted(() => {
  updateTime()
  timeInterval = setInterval(updateTime, 1000)
})

onUnmounted(() => {
  clearInterval(timeInterval)
})
</script>
