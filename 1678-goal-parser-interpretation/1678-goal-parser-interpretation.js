/**
 * @param {string} command
 * @return {string}
 */
var interpret = function(command) {
    let str=""
    let i=0
    while(i<command.length){
        if(command[i]=="G"){
            str+="G"
            i++
        }else if(command[i+1]=="a"){
            str+="al"
            i+=4
        }else{
            str+="o"
            i+=2
        }
    }
    return str
};