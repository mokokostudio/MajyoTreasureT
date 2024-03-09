import { createPVPCavElement, createPVPCav } from '@/animation/cavUtil'

export function Sk1004(_container, _dpr){
  this.skillStatus = 0 // 0 init  1 doing  99 finish
  this.dpr = _dpr
  this.element = createPVPCavElement(_container, _dpr, true, 0)
  this.cav = createPVPCav(this.element, _dpr)
  this.hitelement = createPVPCavElement(_container, _dpr, true, 0)
  this.hitcav = createPVPCav(this.hitelement, _dpr)

  this.do = function(_atker, _beAtker){
    this.skillStatus = 1
    let xOffset = -120
    if(_beAtker.type == 1){
      // 1表示我被攻击
      this.cav.scale(-1, 1);
      this.cav.translate(-this.element.width / this.dpr, 0);
      xOffset = 60
    }
    let diffTime = 0
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/1004/cast_' + idx + '.png')
        img.onload = (() => {
          that.cav.clearRect(0, 0, that.element.width, that.element.height)
          that.cav.drawImage(img, _beAtker.xStand + xOffset, _beAtker.yStand, 200, 95)
        })
        idx++
        if(idx === 115){
          that.skillStatus = 99
          that.element.remove()
          return
        }
        if(idx === 11){
          that.hit(_beAtker)
        }
        idx %= 115
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.hit = function(_beAtker){
    let count = 1
    let diffTime = 0
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/1004/hit_' + idx + '.png')
        img.onload = (() => {
          that.hitcav.clearRect(0, 0, that.hitelement.width, that.hitelement.height)
          that.hitcav.drawImage(img, _beAtker.xStand, _beAtker.yStand - 3, 100, 100)
        })
        idx++
        if(idx === 12){
          count++
          if(count === 7){
            that.hitelement.remove()
            return
          }
        }
        idx %= 12
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

}
