import { createPVPCavElement, createPVPCav } from '@/animation/cavUtil'

export function Sk1007(_container, _dpr){
  this.skillStatus = 0 // 0 init  1 doing  99 finish
  this.dpr = _dpr
  this.element = createPVPCavElement(_container, _dpr, true, 0)
  this.cav = createPVPCav(this.element, _dpr)
  this.hitelement = createPVPCavElement(_container, _dpr, true, 0)
  this.hitcav = createPVPCav(this.hitelement, _dpr)

  this.do = function(_atker, _beAtker){
    this.skillStatus = 1
    this.hitelement.style.zIndex = 20
    let xOffset = -10
    let castXOffset = -10
    if(_beAtker.type == 1){
      // 1表示我被攻击
      this.cav.scale(-1, 1);
      this.cav.translate(-this.element.width / this.dpr, 0);
      this.hitcav.scale(-1, 1);
      this.hitcav.translate(-this.hitelement.width / this.dpr, 0);
      xOffset = 170
      castXOffset = - 190
    }
    let stage = 1
    let maxFrame = 16
    let diffTime = 0
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        if(stage === 1){
          img.src = require('@/assets/frameimg/skills/1007/cast_' + idx + '.png')
        }else if(stage === 2){
          img.src = require('@/assets/frameimg/skills/1007/hit_' + idx + '.png')
        }
        img.onload = (() => {
          that.cav.clearRect(0, 0, that.element.width, that.element.height)
          if(stage === 1){
            that.cav.drawImage(img, _atker.xStand + castXOffset, _atker.yStand - 70, 120, 163)
          }else if(stage === 2){
            that.cav.drawImage(img, _beAtker.xStand + xOffset, _beAtker.yStand - 110, 150, 251)
          }
        })
        if(stage === 2){
          let img2 = new Image()
          img2.src = require('@/assets/frameimg/skills/1007/ground_' + idx + '.png')
          img2.onload = (() => {
            that.hitcav.clearRect(0, 0, that.hitelement.width, that.hitelement.height)
            that.hitcav.drawImage(img2, _beAtker.xStand + xOffset, _beAtker.yStand - 110, 150, 251)
          })
        }
        idx++
        if(idx === maxFrame){
          idx = 0
          if(stage === 1){
            stage = 2
            maxFrame = 33
            diffTime = 34
            xOffset = xOffset - 30
          }else if(stage === 2){
            that.skillStatus = 99
            that.element.remove()
            that.hitelement.remove()
            return
          }
        }
        idx %= maxFrame
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }
}
