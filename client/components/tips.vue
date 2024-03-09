<template>
  <transition name="fade">
    <div class="tip" v-if="isOpen">
      <span class="tip-info">{{ tip }}</span>
    </div>
  </transition>
</template>
<script>
export default {
  data() {
    return {
      tip: '',
      isOpen: false
    }
  },

  methods: {
    openTip(tip, time) {
      this.tip = tip
      this.isOpen = true
      let _time = 2000
      if(time){
        _time = time
      }
      let timer = setTimeout(() => {
        this.isOpen = false
        clearTimeout(timer)
      }, _time)
    }
  }
}

</script>
<style lang="postcss">
  .tip {
    display: flex;
    justify-content: center;
  }

  .tip-info {
    z-index: 9999;
    background-color: rgba(0, 0, 0, 0.9);
    margin-top: 600px;
    padding: 15px;
    border-radius: 12px;
    color: #fff;
    font-size: 30px;
    max-width: 560px;
    white-space: pre-line;
  }


  /** 移除过程 **/
  .fade-enter-active{
    transition: all 0.5s;
  }
  .fade-leave-active {
    transition: all 0.5s;
  }
  .fade-enter{
    opacity: 0;
    transform: translateY(+30px);
  }
  /*** 开始插入、移除结束的位置变化 ***/
  .fade-leave-to {
    opacity: 0;
    transform: translateY(-30px);
  }




</style>
