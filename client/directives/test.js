export function box(a,b){
  this.a = a
  this.b = b
  
  this.ttt = function(){
    let c = 0
    let d = this.b
    console.log(this.a, this.b, c, d, c+this.a)
  }
}