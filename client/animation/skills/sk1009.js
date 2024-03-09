import { createPVPCavElement, createPVPCav } from '@/animation/cavUtil'

export function Sk1009(_container, _dpr){
  this.skillStatus = 0 // 0 init  1 doing  99 finish
  this.dpr = _dpr
  this.element = createPVPCavElement(_container, _dpr, true, 0)
  this.cav = createPVPCav(this.element, _dpr)
  this.hitelement = createPVPCavElement(_container, _dpr, true, 0)
  this.hitcav = createPVPCav(this.hitelement, _dpr)
  this.groundelement = createPVPCavElement(_container, _dpr, true, 0)
  this.groundcav = createPVPCav(this.groundelement, _dpr)

  this.do = function(_atker, _beAtker){
    this.skillStatus = 1
    this.groundelement.style.zIndex = 20
    let xOffset = -16
    let castXOffset = -68
    let castXMove = 0
    let castYMove = 0
    if(_beAtker.type == 1){
      // 1表示我被攻击
      this.cav.scale(-1, 1);
      this.cav.translate(-this.element.width / this.dpr, 0);
      castXOffset = 102
    }
    let diffTime = 20
    let lastDate = Date.now()
    let idx = 0
    let groundIdx = 0
    let showGround = false
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/1009/cast_' + idx + '.png')
        img.onload = (() => {
          that.cav.clearRect(0, 0, that.element.width, that.element.height)
          that.cav.drawImage(img, _beAtker.xStand + castXOffset + castXMove, _beAtker.yStand - 60 + castYMove, 100, 101)
        })
        if(showGround){
          let img2 = new Image()
          img2.src = require('@/assets/frameimg/skills/1009/ground_' + groundIdx + '.png')
          img2.onload = (() => {
            that.groundcav.clearRect(0, 0, that.groundelement.width, that.groundelement.height)
            that.groundcav.drawImage(img2, _beAtker.xStand + xOffset, _beAtker.yStand + 55, 140, 53)
          })
        }
        castXMove += 2
        castYMove += 2
        idx++
        if(showGround){
          groundIdx++
          if(groundIdx === 29){
            that.hit(_beAtker)
          }
          if(groundIdx === 30){
            that.groundelement.remove()
            return
          }
          groundIdx %= 30
        }
        if(idx === 6){
          showGround = true
        }
        if(idx === 22){
          that.element.remove()
        }
        idx %= 22
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.hit = function(_beAtker){
    let diffTime = 34
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/1009/hit_' + idx + '.png')
        img.onload = (() => {
          that.hitcav.clearRect(0, 0, that.hitelement.width, that.hitelement.height)
          that.hitcav.drawImage(img, _beAtker.xStand - 10, _beAtker.yStand + 10, 140, 91)
        })
        idx++
        if(idx === 32){
          that.skillStatus = 99
          that.hitelement.remove()
          return
        }
        idx %= 32
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

}
