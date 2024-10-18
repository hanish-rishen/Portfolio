<template>
  <div class="custom-cursor" ref="cursor">
    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path fill-rule="evenodd" clip-rule="evenodd" d="M13.5 3C13.2239 3 13 3.22386 13 3.5V12C13 12.5523 12.5523 13 12 13C11.4477 13 11 12.5523 11 12V5.5C11 5.22386 10.7761 5 10.5 5C10.2239 5 9.99999 5.22386 9.99999 5.5V13.9677C9.99999 15.0757 8.62655 15.5918 7.8969 14.7579L5.34951 11.8466C5.19167 11.6662 4.95459 11.576 4.71675 11.6057C4.15329 11.6761 3.88804 12.3395 4.24762 12.779L8.93807 18.5118C10.2266 20.0867 12.154 21 14.1888 21C17.3982 21 20 18.3982 20 15.1888V7.5C20 7.22386 19.7761 7 19.5 7C19.2239 7 19 7.22386 19 7.5V12C19 12.5523 18.5523 13 18 13C17.4477 13 17 12.5523 17 12V5.5C17 5.22386 16.7761 5 16.5 5C16.2239 5 16 5.22386 16 5.5V12C16 12.5523 15.5523 13 15 13C14.4477 13 14 12.5523 14 12V3.5C14 3.22386 13.7761 3 13.5 3Z" fill="white"/>
    </svg>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { gsap } from 'gsap'

const cursor = ref(null)
let cursorTimeout

const handleMouseMove = (event) => {
  const { clientX, clientY } = event

  gsap.to(cursor.value, {
    left: clientX,
    top: clientY,
    duration: 0.1,
    ease: 'power2.out'
  })

  gsap.to(cursor.value, {
    opacity: 1,
    duration: 0.3,
    ease: 'power2.out'
  })

  clearTimeout(cursorTimeout)
}

const expandCursor = () => {
  gsap.to(cursor.value, {
    scale: 1.2,
    rotation: -25, // Increased rotation for more tilt
    duration: 0.3,
    ease: 'power2.out'
  })
  cursor.value.classList.add('expanded')
}

const shrinkCursor = () => {
  gsap.to(cursor.value, {
    scale: 1,
    rotation: -10, // Added slight rotation for default state
    duration: 0.3,
    ease: 'power2.out'
  })
  cursor.value.classList.remove('expanded')
}

onMounted(() => {
  window.addEventListener('mousemove', handleMouseMove)
  // Set initial rotation
  gsap.set(cursor.value, { rotation: -10 })
})

onUnmounted(() => {
  window.removeEventListener('mousemove', handleMouseMove)
})
</script>

<style scoped>
.custom-cursor {
  position: fixed;
  width: 32px;
  height: 32px;
  pointer-events: none;
  z-index: 9999;
  transition: transform 0.3s;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  transform: rotate(-10deg); /* Increased default rotation */
}

.custom-cursor svg {
  width: 100%;
  height: 100%;
}

.custom-cursor.expanded {
  transform: scale(1.2) rotate(-25deg); /* Increased rotation for expanded state */
}
</style>