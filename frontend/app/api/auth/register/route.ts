import {NextRequest, NextResponse} from "next/server";

export async function POST(req: NextRequest) {
    const res = await fetch(`${process.env.API_URL}/register`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Cookie: req.headers.get("cookie") || "",
        },
        body: await req.text(),
    });

    const body = await res.text();

    return new NextResponse(body, {
        status: res.status,
    });
}
