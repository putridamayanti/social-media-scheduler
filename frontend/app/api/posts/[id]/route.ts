import {NextRequest, NextResponse} from "next/server";

export async function GET(req: NextRequest, context: { params: Promise<{ id: string }> }) {
    const { id } = await context.params;
    const resp = await fetch(`${process.env.API_URL}/posts/${id}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            Cookie: req.headers.get("cookie") || "",
        },
    });

    const body = await resp.json();

    return NextResponse.json({
        data: body
    })
}

export async function PUT(req: NextRequest, context: { params: Promise<{ id: string }> }) {
    const { id } = await context.params;
    const body = await req.json();
    const resp = await fetch(`${process.env.API_URL}/posts/${id}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
            Cookie: req.headers.get("cookie") || "",
        },
        body: JSON.stringify(body)
    });

    const res = await resp.json();

    return NextResponse.json({
        data: res
    })
}

export async function DELETE(req: NextRequest, context: { params: Promise<{ id: string }> }) {
    const { id } = await context.params;
    const resp = await fetch(`${process.env.API_URL}/posts/${id}`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            Cookie: req.headers.get("cookie") || "",
        },
    });

    const res = await resp.json();

    return NextResponse.json({
        data: res
    })
}
