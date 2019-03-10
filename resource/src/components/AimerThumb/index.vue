<template>
  <div :style="{zIndex:zIndex,height:height,width:width}" class="aimer-item">
    <div class="aimer-info">
      <div class="aimer-info-roles-container">
        <slot/>
      </div>
    </div>
    <img :src="image" class="aimer-thumb">
  </div>
</template>

<script>
export default {
  name: 'AimerThumb',
  props: {
    image: {
      type: String,
      required: true
    },
    zIndex: {
      type: Number,
      default: 1
    },
    width: {
      type: String,
      default: '150px'
    },
    height: {
      type: String,
      default: '150px'
    }
  }
}
</script>

<style scoped>
.aimer-item {
  width: 200px;
  height: 200px;
  border-radius: 50%;
  display: inline-block;
  position: relative;
  cursor: default;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.2);
}

.aimer-info-roles-container {
  padding: 20px;
  text-align: center;
}

.aimer-thumb {
  width: 100%;
  height: 100%;
  background-size: 100%;
  border-radius: 50%;
  overflow: hidden;
  position: absolute;
  transform-origin: 95% 40%;
  transition: all 0.3s ease-in-out;
}

.aimer-thumb:after {
  content: '';
  width: 8px;
  height: 8px;
  position: absolute;
  border-radius: 50%;
  top: 40%;
  left: 95%;
  margin: -4px 0 0 -4px;
  background: radial-gradient(ellipse at center, rgba(14, 14, 14, 1) 0%, rgba(125, 126, 125, 1) 100%);
  box-shadow: 0 0 1px rgba(255, 255, 255, 0.9);
}

.aimer-info {
  position: absolute;
  width: inherit;
  height: inherit;
  border-radius: 50%;
  overflow: hidden;
  box-shadow: inset 0 0 0 5px rgba(0, 0, 0, 0.05);
}

.aimer-info h3 {
  color: #fff;
  text-transform: uppercase;
  position: relative;
  letter-spacing: 2px;
  font-size: 18px;
  margin: 0 60px;
  padding: 22px 0 0 0;
  height: 85px;
  font-family: 'Open Sans', Arial, sans-serif;
  text-shadow: 0 0 1px #fff, 0 1px 2px rgba(0, 0, 0, 0.3);
}

.aimer-info p {
  color: #fff;
  padding: 10px 5px;
  font-style: italic;
  margin: 0 30px;
  font-size: 12px;
  border-top: 1px solid rgba(255, 255, 255, 0.5);
}

.aimer-info p a {
  display: block;
  color: #333;
  width: 80px;
  height: 80px;
  background: rgba(255, 255, 255, 0.3);
  border-radius: 50%;
  color: #fff;
  font-style: normal;
  font-weight: 700;
  text-transform: uppercase;
  font-size: 9px;
  letter-spacing: 1px;
  padding-top: 24px;
  margin: 7px auto 0;
  font-family: 'Open Sans', Arial, sans-serif;
  opacity: 0;
  transition: transform 0.3s ease-in-out 0.2s, opacity 0.3s ease-in-out 0.2s, background 0.2s linear 0s;
  transform: translateX(60px) rotate(90deg);
}

.aimer-info p a:hover {
  background: rgba(255, 255, 255, 0.5);
}

.aimer-item:hover .aimer-thumb {
  transform: rotate(-110deg);
}

.aimer-item:hover .aimer-info p a {
  opacity: 1;
  transform: translateX(0px) rotate(0deg);
}
</style>
