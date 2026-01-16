import DashboardPage from "@/app/dashboard/_components/dashboard-page";
import {cookies} from "next/headers";

export default async function Page() {
    const cookieStore = await cookies();
    const resp = await fetch(`${process.env.NEXT_API_URL}/api/posts`, {
        headers: {
            cookie: cookieStore.toString()
        }
    });
    const res = await resp.json();
    const posts = res?.data?.data;

    return (
        <DashboardPage posts={posts}/>
    )
}
