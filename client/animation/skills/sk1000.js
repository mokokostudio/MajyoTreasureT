import { createPVPCavElement, createPVPCav } from '@/animation/cavUtil'

export function Sk1000(_container, _dpr){
  this.skillStatus = 0 // 0 init  1 doing  99 finish
  this.dpr = _dpr
  this.element = createPVPCavElement(_container, _dpr, true, 0)
  this.cav = createPVPCav(this.element, _dpr)

  this.do = function(_atker, _beAtker){
    this.skillStatus = 1
    let stage = 1
    let maxFrame = 10
    let diffTime = 34
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        if(stage === 1){
          img.src = require('@/assets/frameimg/skills/1000/thd-cast_' + idx + '.png')
        }else if(stage === 2){
          img.src = require('@/assets/frameimg/skills/1000/thd-patk_' + idx + '.png')
        }else if(stage === 3){
          img.src = require('@/assets/frameimg/skills/1000/thd-patk2_' + idx + '.png')
        }else if(stage === 4){
          img.src = require('@/assets/frameimg/skills/1000/thd-atk_' + idx + '.png')
        }
        console.log('thunder stage=+++++++++++',stage,'idx=',idx)
        img.onload = (() => {
          that.cav.clearRect(0, 0, that.element.width, that.element.height)
          if(stage === 1){
            that.cav.drawImage(img, _beAtker.xStand - 50, _beAtker.yStand - 4, 200, 98)
          }else if(stage === 2){
            that.cav.drawImage(img, _beAtker.xStand - 50, _beAtker.yStand - 54, 200, 222)
          }else if(stage === 3){
            that.cav.drawImage(img, _beAtker.xStand - 40, _beAtker.yStand - 54, 180, 185)
          }else if(stage === 4){
            that.cav.drawImage(img, _beAtker.xStand - 67, _beAtker.yStand - 58, 234, 240)
          }
        })
        idx++
        if(idx === maxFrame){
          idx = 0
          if(stage === 1){
            stage = 2
            maxFrame = 14
          }else if(stage === 2){
            stage = 3
            maxFrame = 9
          }else if(stage === 3){
            stage = 4
            maxFrame = 21
          }else if(stage === 4){
            that.skillStatus = 99
            that.element.remove()
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
