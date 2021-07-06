import { EnumType } from "typescript";

export interface User {
        ID:string
        IdString:string
        FirstName: string,
        LastName:string,
        Username: string,
        Password:string,
        Email:string,
        PhoneNumber: string,
        DateOfBirth: Date,
        Gender: EnumType,
        Following: boolean,
        Followers: boolean,
        NumberOfPosts: number,
        WebSite: string
        Biography:string,
        ProfilePicture:string,
        ProfileSettings: {
            Public: boolean,
            MessageRequest: boolean,
            AllowTags: boolean,
            ActiveProfile: boolean
        },
        NotificationSettings: {
            LikeNotification: boolean,
            CommentNotification: boolean,
            MessageRequestNotification: boolean,
            MessageNotification: boolean,
            FollowRequestNotification: boolean,
            FollowNotification: boolean
        },
        VerificationSettings: {
            Verified: boolean,
            "Category": EnumType
        }
}