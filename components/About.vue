<template>
  <section id="about" class="min-h-screen flex flex-col justify-start items-center p-4 relative overflow-hidden" ref="aboutSection">
    <canvas ref="starsCanvas" class="absolute inset-0 z-0"></canvas>
    <div class="flex flex-col md:flex-row w-full max-w-6xl mx-auto mt-24 z-10 relative"> <!-- Added mt-8 class for a small gap above -->
      <div class="w-full md:w-1/2 p-4 flex flex-col justify-center items-center relative">
        <h2 class="title text-4xl md:text-5xl mb-8 font-bold text-white" @mouseenter="expandCursor" @mouseleave="shrinkCursor" ref="aboutTitle">
          About Me
          <span class="underline-animation"></span>
        </h2>
        <div class="profile-container relative mt-4 flex justify-center items-center">
          <img src="/images/profile.jpg" alt="Profile Picture" class="w-64 h-64 object-cover shadow-lg transition-all duration-300" />
          <div class="rotating-border absolute inset-0 border-4 border-dotted border-white"></div>
        </div>
      </div>
      <div class="w-full md:w-1/2 p-4 flex items-center">
        <p class="text-lg md:text-xl text-white" @mouseenter="expandCursor" @mouseleave="shrinkCursor" ref="aboutText">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
        </p>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { gsap } from 'gsap'
import { ScrollTrigger } from 'gsap/ScrollTrigger'
import SplitType from 'split-type'

gsap.registerPlugin(ScrollTrigger)

const aboutSection = ref(null)
const aboutTitle = ref(null)
const aboutText = ref(null)
const starsCanvas = ref(null)

onMounted(() => {
  // GSAP animations
  const timeline = gsap.timeline({
    scrollTrigger: {
      trigger: aboutSection.value,
      start: 'top top',
      end: '+=150%',
      pin: true,
      scrub: 1,
      anticipatePin: 1
    }
  })

  timeline.from(aboutTitle.value, {
    opacity: 0,
    y: -50,
    duration: 0.5
  })

  // Add underline animation
  timeline.from('.underline-animation', {
    width: '0%',
    duration: 0.5,
    ease: 'power1.inOut'
  }, '-=0.3')

  const splitText = new SplitType(aboutText.value, { types: 'words,chars' })
  
  gsap.set(splitText.chars, { opacity: 0.3 })

  timeline.to(splitText.chars, {
    opacity: 1,
    duration: 0.5,
    stagger: 0.02,
    ease: 'power1.inOut'
  })

  ScrollTrigger.create({
    trigger: aboutSection.value,
    start: 'top top',
    end: '+=150%',
    onLeave: () => {
      ScrollTrigger.getById(aboutSection.value)?.kill()
    }
  })

  // Animate rotating border
  gsap.to('.rotating-border', {
    rotation: 360,
    duration: 10,
    repeat: -1,
    ease: 'linear'
  })

  // Initialize star background
  initStarBackground()
})

const expandCursor = () => {
  // Implement cursor expansion logic
}

const shrinkCursor = () => {
  // Implement cursor shrinking logic
}

const initStarBackground = () => {
  const Stars = function(args) {
    if (args === undefined) args = {};
    const _scope = this;

    this.stars = [];
    this.vel = args.vel || 1;
    this.radius = args.radius || 1;
    this.alpha = 0.5;
    this.starsCounter = args.stars || 300;
    const center = {
      x: window.innerWidth / 2,
      y: window.innerHeight / 2
    };
    let canvas, context;

    this.init = function() {
      canvas = starsCanvas.value;
      context = canvas.getContext("2d");
      context.lineCap = "round";
      this.start();
      this.resize();
      window.addEventListener("resize", this.resize.bind(this));
    }

    this.start = function() {
      this.stars = [];
      for (let i = 0; i < this.starsCounter; i++) {
        setTimeout(() => {
          _scope.stars.push(new Star());
        }, i * 30);
      }
    }

    this.resize = function() {
      canvas.width = window.innerWidth;
      canvas.height = window.innerHeight;
      center.x = canvas.width / 2;
      center.y = canvas.height / 2;
    }

    this.animate = function() {
      window.requestAnimationFrame(this.animate.bind(this));
      this.render();
    }

    this.render = function() {
      context.clearRect(0, 0, canvas.width, canvas.height);
      context.fillStyle = 'rgba(0, 0, 0, 0.8)';
      context.fillRect(0, 0, canvas.width, canvas.height);
      context.strokeStyle = "white";
      for (let i = 0; i < this.stars.length; i++) this.stars[i].update();
    }

    const Star = function() {
      this.x = center.x;
      this.y = center.y;
      this.init = function() {
        this.radius = Math.random() * _scope.radius;
        this.x = center.x;
        this.y = center.y;
        this.lineWidth = 0;
        this.vel = {
          x: Math.random() * 10 - 5,
          y: Math.random() * 10 - 5
        }
      }
      this.update = function() {
        this.vel.x *= 1.05;
        this.vel.y *= 1.05;
        this.lineWidth += 0.035;
        this.x0 = this.x;
        this.y0 = this.y;
        this.x += this.vel.x;
        this.y += this.vel.y;
        this.draw();
        if (this.isDead()) this.init();
      }
      this.draw = function() {
        context.beginPath();
        context.moveTo(this.x0, this.y0);
        context.lineTo(this.x, this.y);
        context.lineWidth = this.lineWidth;
        context.stroke();
      }
      this.isDead = function() {
        return (this.x < 0 || this.x > canvas.width || this.y < 0 || this.y > canvas.height);
      }
      this.init();
      return this;
    }
    this.init();
    this.animate();
    return this;
  }

  new Stars();
}
</script>

<style scoped>
.profile-container {
  transition: all 0.3s ease;
  position: relative;
  width: 256px;
  height: 256px;
  margin: 0 auto;
}

.rotating-border {
  transition: all 0.3s ease;
  width: 100%;
  height: 100%;
  box-sizing: border-box;
  pointer-events: none;
}

.title {
  position: relative;
  display: inline-block;
}

.underline-animation {
  position: absolute;
  bottom: -5px;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: currentColor;
}

body {
  overflow-x: hidden;
}
</style>