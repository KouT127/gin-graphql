import React, {useState} from "react";
import * as firebase from "../../config/firebase";
import * as FirebaseType from "firebase"

interface User {
    id: string,
    name: string | null,
    email: string | null,
    emailVerified: boolean,
    token: string | null,
}

interface UserState {
    user: User | null
}

type UseUserType = [UserState, React.Dispatch<React.SetStateAction<UserState>>] | null;

const initialUserData: UserState = {user: null};


const UserContext = React.createContext<UseUserType>(null);

// useStateをProvideする。
const UserProvider = (props: any) => {
    const value = useState<UserState>(initialUserData);
    return <UserContext.Provider value={value} {...props} />
};

//　useUserStateを使用した階層からUserContextを辿り、
//　存在した場合、上位のuseStateを取得する。
//　setUserStateは流さず、変更はここに閉じ込めておく。
const useUserState = () => {
    const context = React.useContext(UserContext);
    if (!context) {
        throw new Error(`Not Provide UserContext`)
    }
    const [userState, setUserState] = context;

    const isLoggedIn = () => {
        if (!userState) return null;
        if (!userState.user) return false;
        return true
    };

    const authConnect = () => {
        firebase.default.auth().onAuthStateChanged(async (user: FirebaseType.User | null) => {
                console.log('Auth');
                if (!user) {
                    return;
                }
                console.log('After Auth');
                const authUser: User = {
                    id: user.uid,
                    name: user.displayName,
                    email: user.email,
                    emailVerified: user.emailVerified,
                    token: await user.getIdToken()
                };
                const userData: UserState = {user: authUser};
                setUserState && setUserState(userData)
            }
        )
    };

    const signIn = (email: string, password: string) => {
        console.log('In');
        firebase.default.auth().signInWithEmailAndPassword(email, password).then(async (result: FirebaseType.auth.UserCredential | null) => {
            const user = result!.user;
            if (!user) {
                return;
            }
            const authUser: User = {
                id: user.uid,
                name: user.displayName,
                email: user.email,
                emailVerified: user.emailVerified,
                token: await user.getIdToken()
            };
            const userData: UserState = {user: authUser};
            setUserState && setUserState(userData)
        }).catch((e) => {
            console.log(e)
        })
    };

    const signOut = () => {
        console.log('Out');
        firebase.default.auth().signOut().then(() => {
            setUserState({user: null})
        })
    };

    return {
        userState,
        isLoggedIn,
        authConnect,
        signIn,
        signOut,
    }
};


export {UserProvider, useUserState};