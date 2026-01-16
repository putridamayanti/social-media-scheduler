import {Button} from "@/components/ui/button";
import {useRouter} from "next/navigation";

export default function Navbar() {
    const router = useRouter();

    const handleLogout = async () => {
        const resp = await fetch('/api/auth/logout', {
            method: 'DELETE',
        });
        if (resp.status === 200) {
            router.push('/');
        }
    };

    return (
        <header className="container flex justify-between items-center">
            <h1 className="max-w-xs text-3xl font-semibold leading-10 tracking-tight text-black dark:text-zinc-50">
                Social Media Post Scheduler
            </h1>

            <div className="flex gap-6">
                <Button onClick={handleLogout}>
                    Logout
                </Button>
            </div>
        </header>
    )
}
