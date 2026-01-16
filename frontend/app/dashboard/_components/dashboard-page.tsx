'use client'

import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card";
import {Item, ItemActions, ItemContent, ItemTitle} from "@/components/ui/item";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import {Button} from "@/components/ui/button";
import {convertDateTimeToLocal} from "@/lib/helpers";
import {Badge} from "@/components/ui/badge";
import {useMemo, useState} from "react";
import {Dialog, DialogContent, DialogHeader, DialogTitle} from "@/components/ui/dialog";
import PostForm from "@/app/dashboard/_components/post-form";
import {Pencil, Trash} from "lucide-react";
import {useRouter} from "next/navigation";
import Navbar from "@/components/navbar";

interface DashboardPageType {
    posts: any[]
}

export default function DashboardPage(props: DashboardPageType) {
    const { posts } = props;
    const router = useRouter();

    const [formDialog, setFormDialog] = useState({open: false, data: null});

    const items = useMemo(() => {
        return {
            upcoming: posts.filter(item => item.status !== 'published'),
            history: posts.filter(item => item.status === 'published')
        }
    }, [posts])

    const handleDelete = async (id: string) => {
        const resp = await fetch(`/api/posts/${id}`, {
            method: 'DELETE',
        });
        if (resp.status === 200) {
            router.refresh();
        }
    }

    return (
        <>
            <Navbar/>
            <main className="min-h-screen flex justify-center items-center bg-muted">
                <div className="flex w-full max-w-lg flex-col gap-6">
                    <Button onClick={() => setFormDialog({open: true, data: null})}>
                        Add Post
                    </Button>
                    <Tabs defaultValue="upcoming">
                        <TabsList>
                            <TabsTrigger value="upcoming">Upcoming</TabsTrigger>
                            <TabsTrigger value="history">History</TabsTrigger>
                        </TabsList>
                        <TabsContent value="upcoming">
                            <Card>
                                <CardHeader>
                                    <CardTitle>Upcoming</CardTitle>
                                    <CardDescription>
                                        List of upcoming posts
                                    </CardDescription>
                                </CardHeader>
                                <CardContent className="grid gap-6">
                                    {items.upcoming.map((item, i) => (
                                        <Item key={i} variant="outline">
                                            <ItemContent>
                                                <ItemTitle>{item.title}</ItemTitle>
                                                <div>
                                                    <p>Scheduled At: <span className="font-semibold">{convertDateTimeToLocal(item.scheduled_at)}</span></p>
                                                </div>
                                                <Badge variant={item.status === 'published' ? 'default' : 'secondary'} className="mt-4 capitalize">
                                                    {item.status}
                                                </Badge>
                                            </ItemContent>
                                            <ItemActions>
                                                <Button variant="outline" size="icon-sm" onClick={() => setFormDialog({open: true, data: item})}>
                                                    <Pencil/>
                                                </Button>
                                                <Button variant="destructive" size="icon-sm" onClick={() => handleDelete(item.id)}>
                                                    <Trash/>
                                                </Button>
                                            </ItemActions>
                                        </Item>
                                    ))}
                                </CardContent>
                            </Card>
                        </TabsContent>
                        <TabsContent value="history">
                            <Card>
                                <CardHeader>
                                    <CardTitle>History</CardTitle>
                                    <CardDescription>
                                        List of published posts
                                    </CardDescription>
                                </CardHeader>
                                <CardContent className="grid gap-6">
                                    {items.history.map((item, i) => (
                                        <Item key={i} variant="outline">
                                            <ItemContent>
                                                <ItemTitle>{item.title}</ItemTitle>
                                                <div>
                                                    <p>Scheduled At: <span className="font-semibold">{convertDateTimeToLocal(item.scheduled_at)}</span></p>
                                                    <p>Published At: <span className="font-semibold">{convertDateTimeToLocal(item.published_at)}</span></p>
                                                </div>
                                                <Badge variant={item.status === 'published' ? 'default' : 'secondary'} className="mt-4 capitalize">
                                                    {item.status}
                                                </Badge>
                                            </ItemContent>
                                        </Item>
                                    ))}
                                </CardContent>
                            </Card>
                        </TabsContent>
                    </Tabs>
                </div>

                <Dialog open={formDialog.open} onOpenChange={() => setFormDialog({open: !formDialog.open, data: null})}>
                    <DialogContent>
                        <DialogHeader>
                            <DialogTitle>{formDialog.data ? 'Update' : 'Create'} Post</DialogTitle>

                            <PostForm data={formDialog.data} onClose={() => setFormDialog({open: false, data: null})}/>
                        </DialogHeader>
                    </DialogContent>
                </Dialog>
            </main>
        </>
    )
}
