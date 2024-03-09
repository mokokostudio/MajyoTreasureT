import { createPVPCavElement, createPVPCav } from '@/animation/cavUtil'

export function Sk1005(_container, _dpr){
  this.skillStatus = 0 // 0 init  1 doing  99 finish
  this.dpr = _dpr
  this.element = createPVPCavElement(_container, _dpr, true, 0)
  this.cav = createPVPCav(this.element, _dpr)

  this.do = function(_atker, _beAtker){
    this.skillStatus = 1
    let diffTime = 20
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/1005/hit_' + idx + '.png')
        img.onload = (() => {
          that.cav.clearRect(0, 0, that.element.width, that.element.height)
          that.cav.drawImage(img, _beAtker.xStand - 20, _beAtker.yStand - 80, 140, 197)
        })
        idx++
        if(idx === 45){
          that.skillStatus = 99
          that.element.remove()
          return
        }
        idx %= 45
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

}
