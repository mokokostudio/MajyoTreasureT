export function Skill(_baseCav, _baseElement, _dpr){
  this.baseCav = _baseCav
  this.baseElement = _baseElement
  this.dpr = _dpr
  this.skillStatus = 0 // 0 初始化 99 技能结束  1 进行中

  this.scaleX = function(){
    this.baseCav.scale(-1, 1);
    this.baseCav.translate(-this.baseElement.width / this.dpr, 0);
  }

  this.rescaleX = function(){
    this.baseCav.scale(-1, 1);
    this.baseCav.translate(-this.baseElement.width / this.dpr, 0);
  }

  this.thunder = function(_character){
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
          img.src = require('@/assets/frameimg/skills/thunder/thd-cast_' + idx + '.png')
        }else if(stage === 2){
          img.src = require('@/assets/frameimg/skills/thunder/thd-patk_' + idx + '.png')
        }else if(stage === 3){
          img.src = require('@/assets/frameimg/skills/thunder/thd-patk2_' + idx + '.png')
        }else if(stage === 4){
          img.src = require('@/assets/frameimg/skills/thunder/thd-atk_' + idx + '.png')
        }
        // console.log('stage=',stage,'idx=',idx)
        img.onload = (() => {
          that.baseCav.clearRect(0, 0, that.baseElement.width, that.baseElement.height)
          if(stage === 1){
            that.baseCav.drawImage(img, _character.xStand - 50, _character.yStand - 4, 200, 98)
          }else if(stage === 2){
            that.baseCav.drawImage(img, _character.xStand - 50, _character.yStand - 54, 200, 222)
          }else if(stage === 3){
            that.baseCav.drawImage(img, _character.xStand - 40, _character.yStand - 54, 180, 185)
          }else if(stage === 4){
            that.baseCav.drawImage(img, _character.xStand - 67, _character.yStand - 58, 234, 240)
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
            return
          }
        }
        idx %= maxFrame
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk102101 = function(_character, _hitCav, _hitElement){
    this.skillStatus = 1
    let xOffset = -120
    if(_character.type == 1){
      // 1表示我被攻击
      this.scaleX()
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
        img.src = require('@/assets/frameimg/skills/102101/cast_' + idx + '.png')
        img.onload = (() => {
          that.baseCav.clearRect(0, 0, that.baseElement.width, that.baseElement.height)
          that.baseCav.drawImage(img, _character.xStand + xOffset, _character.yStand, 200, 95)
        })
        idx++
        if(idx === 115){
          that.skillStatus = 99
          if(_character.type == 1){
            that.rescaleX()
          }
          return
        }
        if(idx === 11){
          that.sk102101Hit(_character, _hitCav, _hitElement)
        }
        idx %= 115
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk102101Hit = function(_character, _hitCav, _hitElement){
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
        img.src = require('@/assets/frameimg/skills/102101/hit_' + idx + '.png')
        img.onload = (() => {
          _hitCav.clearRect(0, 0, _hitElement.width, _hitElement.height)
          _hitCav.drawImage(img, _character.xStand, _character.yStand - 3, 100, 100)
        })
        idx++
        if(idx === 12){
          count++
          if(count === 7){
            _hitElement.remove()
            return
          }
        }
        idx %= 12
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk102001 = function(_character,_castCav, _castElement, _hitCav, _hitElement){
    this.skillStatus = 1
    let xOffset = -150
    let count = 0
    if(_character.type == 1){
      // 1表示我被攻击
      _castCav.scale(-1, 1);
      _castCav.translate(-_castElement.width / this.dpr, 0);
      xOffset = 30
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
        img.src = require('@/assets/frameimg/skills/102001/cast_' + idx + '.png')
        img.onload = (() => {
          _castCav.clearRect(0, 0, _castElement.width, _castElement.height)
          _castCav.drawImage(img, _character.xStand + xOffset, _character.yStand + 7, 143, 80)
        })
        idx++
        xOffset = xOffset + 8
        if(idx === 15){
          count++
          if(count === 1){
            that.sk102001Hit(_character, _hitCav, _hitElement)
          }
          if(count === 3){
            that.skillStatus = 99
            _castElement.remove()
            return
          }
        }
        idx %= 15
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk102001Hit = function(_character, _hitCav, _hitElement){
    let diffTime = 0
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/102001/hit_' + idx + '.png')
        img.onload = (() => {
          _hitCav.clearRect(0, 0, _hitElement.width, _hitElement.height)
          _hitCav.drawImage(img, _character.xStand, _character.yStand - 2.5, 100, 101)
        })
        idx++
        if(idx === 14){
          _hitElement.remove()
          return
        }
        idx %= 14
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk100901 = function(_character, _castCav, _castElement, _hitCav, _hitElement){
    this.skillStatus = 1
    let xOffset = -2
    if(_character.type == 1){
      // 1表示我被攻击
      _castCav.scale(-1, 1);
      _castCav.translate(-_castElement.width / this.dpr, 0);
      xOffset = 180
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
        img.src = require('@/assets/frameimg/skills/100901/cast_' + idx + '.png')
        img.onload = (() => {
          _castCav.clearRect(0, 0, _castElement.width, _castElement.height)
          _castCav.drawImage(img, _character.xStand + xOffset, _character.yStand + 44, 87, 50)
        })
        idx++
        if(idx === 35){
          that.sk100901Hit(_character, _hitCav, _hitElement)
          _castElement.remove()
          return
        }
        idx %= 35
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk100901Hit = function(_character, _hitCav, _hitElement){
    let diffTime = 0
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/100901/hit_' + idx + '.png')
        img.onload = (() => {
          _hitCav.clearRect(0, 0, _hitElement.width, _hitElement.height)
          _hitCav.drawImage(img, _character.xStand - 5, _character.yStand - 5, 110, 105)
        })
        idx++
        if(idx === 40){
          that.skillStatus = 99
          _hitElement.remove()
          return
        }
        idx %= 40
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk102301 = function(_character, _hitCav, _hitElement){
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
        img.src = require('@/assets/frameimg/skills/102301/hit_' + idx + '.png')
        img.onload = (() => {
          _hitCav.clearRect(0, 0, _hitElement.width, _hitElement.height)
          _hitCav.drawImage(img, _character.xStand - 20, _character.yStand - 80, 140, 197)
        })
        idx++
        if(idx === 45){
          that.skillStatus = 99
          _hitElement.remove()
          return
        }
        idx %= 45
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk103501 = function(_character, _castCav, _castElement, _hitCav, _hitElement){
    this.skillStatus = 1
    let xOffset = -10
    if(_character.type == 1){
      // 1表示我被攻击
      _castCav.scale(-1, 1);
      _castCav.translate(-_castElement.width / this.dpr, 0);
      xOffset = 170
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
        img.src = require('@/assets/frameimg/skills/103501/cast_' + idx + '.png')
        img.onload = (() => {
          _castCav.clearRect(0, 0, _castElement.width, _castElement.height)
          _castCav.drawImage(img, _character.xStand + xOffset, _character.yStand - 20, 100, 139)
        })
        idx++
        if(idx === 46){
          that.sk103501Hit(_character, _hitCav, _hitElement)
          _castElement.remove()
          return
        }
        idx %= 46
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk103501Hit = function(_character, _hitCav, _hitElement){
    let diffTime = 34
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/103501/hit_' + idx + '.png')
        img.onload = (() => {
          _hitCav.clearRect(0, 0, _hitElement.width, _hitElement.height)
          _hitCav.drawImage(img, _character.xStand - 30, _character.yStand - 30, 160, 134)
        })
        idx++
        if(idx === 24){
          that.skillStatus = 99
          _hitElement.remove()
          return
        }
        idx %= 24
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk104301 = function(_character, _characterMy, _castCav, _castElement, _groundCav, _groundElement){
    this.skillStatus = 1
    _groundElement.style.zIndex = 20
    let xOffset = -10
    let castXOffset = -10
    if(_character.type == 1){
      // 1表示我被攻击
      _castCav.scale(-1, 1);
      _castCav.translate(-_castElement.width / this.dpr, 0);
      _groundCav.scale(-1, 1);
      _groundCav.translate(-_groundElement.width / this.dpr, 0);
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
          img.src = require('@/assets/frameimg/skills/104301/cast_' + idx + '.png')
        }else if(stage === 2){
          img.src = require('@/assets/frameimg/skills/104301/hit_' + idx + '.png')
        }
        img.onload = (() => {
          _castCav.clearRect(0, 0, _castElement.width, _castElement.height)
          if(stage === 1){
            _castCav.drawImage(img, _characterMy.xStand + castXOffset, _characterMy.yStand - 70, 120, 163)
          }else if(stage === 2){
            _castCav.drawImage(img, _character.xStand + xOffset, _character.yStand - 110, 150, 251)
          }
        })
        if(stage === 2){
          let img2 = new Image()
          img2.src = require('@/assets/frameimg/skills/104301/ground_' + idx + '.png')
          img2.onload = (() => {
            _groundCav.clearRect(0, 0, _groundElement.width, _groundElement.height)
            _groundCav.drawImage(img2, _character.xStand + xOffset, _character.yStand - 110, 150, 251)
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
            _castElement.remove()
            _groundElement.remove()
            return
          }
        }
        idx %= maxFrame
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk104801 = function(_character, _castCav, _castElement){
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
        img.src = require('@/assets/frameimg/skills/104801/cast_' + idx + '.png')
        img.onload = (() => {
          _castCav.clearRect(0, 0, _castElement.width, _castElement.height)
          _castCav.drawImage(img, _character.xStand - 20, _character.yStand - 40, 140, 170)
        })
        idx++
        if(idx === 55){
          that.skillStatus = 99
          _castElement.remove()
          return
        }
        idx %= 55
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk105001 = function(_character, _castCav, _castElement, _groundCav, _groundElement, _hitCav, _hitElement){
    this.skillStatus = 1
    _groundElement.style.zIndex = 20
    let xOffset = -16
    let castXOffset = -68
    let castXMove = 0
    let castYMove = 0
    if(_character.type == 1){
      // 1表示我被攻击
      _castCav.scale(-1, 1);
      _castCav.translate(-_castElement.width / this.dpr, 0);
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
        img.src = require('@/assets/frameimg/skills/105001/cast_' + idx + '.png')
        img.onload = (() => {
          _castCav.clearRect(0, 0, _castElement.width, _castElement.height)
          _castCav.drawImage(img, _character.xStand + castXOffset + castXMove, _character.yStand - 60 + castYMove, 100, 101)
        })
        if(showGround){
          let img2 = new Image()
          img2.src = require('@/assets/frameimg/skills/105001/ground_' + groundIdx + '.png')
          img2.onload = (() => {
            _groundCav.clearRect(0, 0, _groundElement.width, _groundElement.height)
            _groundCav.drawImage(img2, _character.xStand + xOffset, _character.yStand + 55, 140, 53)
          })
        }
        castXMove += 2
        castYMove += 2
        idx++
        if(showGround){
          groundIdx++
          if(groundIdx === 29){
            that.sk105001Hit(_character, _hitCav, _hitElement)
          }
          if(groundIdx === 30){
            _groundElement.remove()
            return
          }
          groundIdx %= 30
        }
        if(idx === 6){
          showGround = true
        }
        if(idx === 22){
          _castElement.remove()
        }
        idx %= 22
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk105001Hit = function(_character, _hitCav, _hitElement){
    let diffTime = 34
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/105001/hit_' + idx + '.png')
        img.onload = (() => {
          _hitCav.clearRect(0, 0, _hitElement.width, _hitElement.height)
          _hitCav.drawImage(img, _character.xStand - 10, _character.yStand + 10, 140, 91)
        })
        idx++
        if(idx === 32){
          that.skillStatus = 99
          _hitElement.remove()
          return
        }
        idx %= 32
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk101001 = function(_character, _castCav, _castElement, _hitCav, _hitElement){
    this.skillStatus = 1
    let diffTime = 0
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/101001/cast_' + idx + '.png')
        img.onload = (() => {
          _castCav.clearRect(0, 0, _castElement.width, _castElement.height)
          _castCav.drawImage(img, _character.xStand, _character.yStand - 30, 100, 166)
        })
        idx++
        if(idx === 33){
          that.sk101001Hit(_character, _hitCav, _hitElement)
        }
        if(idx === 35){
          _castElement.remove()
          return
        }
        idx %= 35
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk101001Hit = function(_character, _hitCav, _hitElement){
    let diffTime = 20
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/101001/hit_' + idx + '.png')
        img.onload = (() => {
          _hitCav.clearRect(0, 0, _hitElement.width, _hitElement.height)
          _hitCav.drawImage(img, _character.xStand - 20, _character.yStand - 66, 140, 180)
        })
        idx++
        if(idx === 41){
          that.skillStatus = 99
          _hitElement.remove()
          return
        }
        idx %= 41
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk105301 = function(_character, _hitCav, _hitElement){
    this.skillStatus = 1
    let diffTime = 34
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/105301/cast_' + idx + '.png')
        img.onload = (() => {
          _hitCav.clearRect(0, 0, _hitElement.width, _hitElement.height)
          _hitCav.drawImage(img, _character.xStand - 40, _character.yStand - 100, 200, 253)
        })
        idx++
        if(idx === 49){
          that.skillStatus = 99
          _hitElement.remove()
          return
        }
        idx %= 49
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.sk105302 = function(_character, _hitCav, _hitElement){
    this.skillStatus = 1
    let diffTime = 34
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/105302/cast_' + idx + '.png')
        img.onload = (() => {
          _hitCav.clearRect(0, 0, _hitElement.width, _hitElement.height)
          _hitCav.drawImage(img, _character.xStand - 30, _character.yStand - 60, 160, 177)
        })
        idx++
        if(idx === 52){
          that.skillStatus = 99
          _hitElement.remove()
          return
        }
        idx %= 52
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }
}
