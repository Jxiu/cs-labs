import { reactive } from 'vue'
import axios from 'axios';

export default () =>{
    // https://dog.ceo/api/breed/pembroke/images/random
    let dogList = reactive([
        "https://images.dog.ceo/breeds/dhole/n02115913_25.jpg"
    ])
    
    async function moreDog(){
        try{            
            let result = await axios.get("https://dog.ceo/api/breed/pembroke/images/random")
            console.log(result.data)
            dogList.push(result.data.message)
        }catch(err){
            console.log(err)
        }
    }
    // 注意： 必须要有返回值
    return {dogList, moreDog}
}