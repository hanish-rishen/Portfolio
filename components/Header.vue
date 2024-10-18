<template>
  <header class="min-h-screen flex flex-col items-center justify-center relative overflow-hidden py-8 px-4 md:px-12 lg:px-24">
    <div class="net-background absolute inset-0" ref="netBackground"></div>
    <div class="scifi-objects absolute inset-0 z-0" ref="scifiObjects"></div>
    <div class="absolute top-4 left-4 bg-white bg-opacity-10 backdrop-blur-sm px-4 py-2 rounded" ref="welcomeText">
      <span class="text-white text-lg font-medium">Welcome!</span>
    </div>
    <div class="text-center relative z-10">
      <h1 class="title text-6xl md:text-8xl lg:text-9xl mb-4" @mouseenter="expandCursor" @mouseleave="shrinkCursor" ref="nameText"></h1>
      <p class="subtitle text-3xl md:text-4xl lg:text-5xl" @mouseenter="expandCursor" @mouseleave="shrinkCursor" ref="roleText"></p>
    </div>
    <div class="absolute bottom-4 right-4 bg-white bg-opacity-10 backdrop-blur-sm px-4 py-2 rounded" ref="scrollDownText">
      <span class="text-white text-lg font-medium">Scroll Down ↓</span>
    </div>
  </header>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { gsap } from 'gsap'
import * as THREE from 'three'

const netBackground = ref(null)
const scifiObjects = ref(null)
const welcomeText = ref(null)
const scrollDownText = ref(null)
const nameText = ref(null)
const roleText = ref(null)

let scene, camera, renderer

const initSciFiObjects = () => {
  scene = new THREE.Scene()
  camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000)
  renderer = new THREE.WebGLRenderer({ alpha: true })
  renderer.setSize(window.innerWidth, window.innerHeight)
  scifiObjects.value.appendChild(renderer.domElement)

  // Create minimalistic sci-fi objects
  const geometry1 = new THREE.TorusGeometry(5, 2, 16, 100)
  const material1 = new THREE.MeshBasicMaterial({ color: 0x00ffff, wireframe: true })
  const torus = new THREE.Mesh(geometry1, material1)
  torus.position.set(15, 15, -10)
  scene.add(torus)

  const geometry2 = new THREE.OctahedronGeometry(4)
  const material2 = new THREE.MeshBasicMaterial({ color: 0xff00ff, wireframe: true })
  const octahedron = new THREE.Mesh(geometry2, material2)
  octahedron.position.set(15, -10, -5)
  scene.add(octahedron)

  const geometry3 = new THREE.IcosahedronGeometry(3)
  const material3 = new THREE.MeshBasicMaterial({ color: 0xffff00, wireframe: true })
  const icosahedron = new THREE.Mesh(geometry3, material3)
  icosahedron.position.set(0, 15, -15)
  scene.add(icosahedron)

  const geometry4 = new THREE.DodecahedronGeometry(3)
  const material4 = new THREE.MeshBasicMaterial({ color: 0xff6600, wireframe: true })
  const dodecahedron = new THREE.Mesh(geometry4, material4)
  dodecahedron.position.set(-20, -15, -8)
  scene.add(dodecahedron)

  const geometry6 = new THREE.ConeGeometry(2, 4, 32)
  const material6 = new THREE.MeshBasicMaterial({ color: 0x0066ff, wireframe: true })
  const cone = new THREE.Mesh(geometry6, material6)
  cone.position.set(-18, 18, -10)
  scene.add(cone)

  const geometry7 = new THREE.CylinderGeometry(1, 2, 4, 32)
  const material7 = new THREE.MeshBasicMaterial({ color: 0xff0066, wireframe: true })
  const cylinder = new THREE.Mesh(geometry7, material7)
  cylinder.position.set(18, -18, -6)
  scene.add(cylinder)

  const geometry9 = new THREE.BoxGeometry(3, 3, 3)
  const material9 = new THREE.MeshBasicMaterial({ color: 0x6600ff, wireframe: true })
  const cube = new THREE.Mesh(geometry9, material9)
  cube.position.set(0, -15, -10)
  scene.add(cube)

  const geometry10 = new THREE.RingGeometry(2, 3, 32)
  const material10 = new THREE.MeshBasicMaterial({ color: 0xff6600, wireframe: true, side: THREE.DoubleSide })
  const ring = new THREE.Mesh(geometry10, material10)
  ring.position.set(-10, 10, -12)
  scene.add(ring)

  camera.position.z = 30

  const animate = () => {
    requestAnimationFrame(animate)
    torus.rotation.x += 0.01
    torus.rotation.y += 0.01
    octahedron.rotation.x -= 0.01
    octahedron.rotation.z -= 0.01
    icosahedron.rotation.y += 0.01
    icosahedron.rotation.z += 0.01
    dodecahedron.rotation.x += 0.01
    dodecahedron.rotation.y -= 0.01
    cone.rotation.y += 0.01
    cylinder.rotation.z -= 0.01
    cube.rotation.x += 0.01
    cube.rotation.y += 0.01
    ring.rotation.x += 0.01
    ring.rotation.y += 0.01
    renderer.render(scene, camera)
  }

  animate()
}

const scrambleText = (element, finalText) => {
  const chars = '!<>-_\\/[]{}—=+*^?#________'
  let iteration = 0
  const finalDuration = 1500

  const interval = setInterval(() => {
    element.textContent = finalText
      .split('')
      .map((letter, index) => {
        if (index < iteration) {
          return finalText[index]
        }
        return chars[Math.floor(Math.random() * chars.length)]
      })
      .join('')

    if (iteration >= finalText.length) {
      clearInterval(interval)
    }

    iteration += 1 / 3
  }, finalDuration / finalText.length)
}

onMounted(() => {
  const canvas = document.createElement('canvas')
  const ctx = canvas.getContext('2d')
  canvas.width = window.innerWidth
  canvas.height = window.innerHeight
  netBackground.value.appendChild(canvas)

  const drawNet = () => {
    ctx.strokeStyle = 'rgba(255, 255, 255, 0.2)'
    ctx.lineWidth = 1.5

    const cellSize = 40
    for (let x = 0; x < canvas.width; x += cellSize) {
      ctx.beginPath()
      ctx.moveTo(x, 0)
      ctx.lineTo(x, canvas.height)
      ctx.stroke()
    }

    for (let y = 0; y < canvas.height; y += cellSize) {
      ctx.beginPath()
      ctx.moveTo(0, y)
      ctx.lineTo(canvas.width, y)
      ctx.stroke()
    }
  }

  drawNet()

  gsap.from(welcomeText.value, {
    opacity: 0,
    y: -50,
    duration: 1,
    ease: 'power3.out'
  })

  gsap.from(scrollDownText.value, {
    opacity: 0,
    y: 50,
    duration: 1,
    ease: 'power3.out',
    delay: 0.5
  })

  initSciFiObjects()

  setTimeout(() => {
    scrambleText(nameText.value, 'Hanish Rishen')
  }, 500)

  setTimeout(() => {
    scrambleText(roleText.value, 'Web Developer & Designer')
  }, 1000)
})

const expandCursor = () => {
  // Implement cursor expansion logic
}

const shrinkCursor = () => {
  // Implement cursor shrinking logic
}
</script>