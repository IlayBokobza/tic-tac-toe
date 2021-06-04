import { createSlice } from '@reduxjs/toolkit'

export const storeSlice = createSlice({
    name:'store',
    initialState:{
        socket:null,
    },
    reducers:{
        setSocket(state:any,socket){state.socket = socket},
    },
})

export const {setSocket} = storeSlice.actions
export default storeSlice.reducer