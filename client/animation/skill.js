import { Sk1000 } from '@/animation/skills/sk1000'
import { Sk1001 } from '@/animation/skills/sk1001'
import { Sk1002 } from '@/animation/skills/sk1002'
import { Sk1003 } from '@/animation/skills/sk1003'
import { Sk1004 } from '@/animation/skills/sk1004'
import { Sk1005 } from '@/animation/skills/sk1005'
import { Sk1006 } from '@/animation/skills/sk1006'
import { Sk1007 } from '@/animation/skills/sk1007'
import { Sk1008 } from '@/animation/skills/sk1008'
import { Sk1009 } from '@/animation/skills/sk1009'
import { Sk1010 } from '@/animation/skills/sk1010'
import { Sk1011 } from '@/animation/skills/sk1011'

export function Skill(_container, _dpr){
  this.ref = _container
  this.dpr = _dpr
  this.status = 0

  this.cast = function(_skillCode, _atker, _beAtker){
    this.status = 1
    let sk = null
    if(_skillCode === "1000"){
      sk = new Sk1000(this.ref, this.dpr)
    }else if(_skillCode === "1001"){
      sk = new Sk1001(this.ref, this.dpr)
    }else if(_skillCode === "1002"){
      sk = new Sk1002(this.ref, this.dpr)
    }else if(_skillCode === "1003"){
      sk = new Sk1003(this.ref, this.dpr)
    }else if(_skillCode === "1004"){
      sk = new Sk1004(this.ref, this.dpr)
    }else if(_skillCode === "1005"){
      sk = new Sk1005(this.ref, this.dpr)
    }else if(_skillCode === "1006"){
      sk = new Sk1006(this.ref, this.dpr)
    }else if(_skillCode === "1007"){
      sk = new Sk1007(this.ref, this.dpr)
    }else if(_skillCode === "1008"){
      sk = new Sk1008(this.ref, this.dpr)
    }else if(_skillCode === "1009"){
      sk = new Sk1009(this.ref, this.dpr)
    }else if(_skillCode === "1010"){
      sk = new Sk1010(this.ref, this.dpr)
    }else if(_skillCode === "1011"){
      sk = new Sk1011(this.ref, this.dpr)
    }
    sk.do(_atker, _beAtker)
    this.finish(sk, _atker)
  }

  this.finish = function(_skillObj, _atker){
    let sktime = setInterval(() => {
      if(_skillObj.skillStatus == 99){
        this.status = 99
        _skillObj = null
        clearInterval(sktime)
      }
    },30)
  }
}
