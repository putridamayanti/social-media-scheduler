import {NextRequest, NextResponse} from "next/server";

export async function GET(req: NextRequest) {
    const resp = await fetch(`${process.env.API_URL}/posts`, {
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

export async function POST(req: NextRequest) {
    const body = await req.json();
    const resp = await fetch(`${process.env.API_URL}/posts`, {
        method: "POST",
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
