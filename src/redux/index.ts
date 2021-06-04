import {configureStore, getDefaultMiddleware} from '@reduxjs/toolkit'
import storeReducer from './store'

export const store = configureStore({
    reducer:{
        storeReducer
    },
    middleware:getDefaultMiddleware({
        serializableCheck:{
            ignoredActions:['store/setSocket']   
        }
    })
},)

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch

export default store