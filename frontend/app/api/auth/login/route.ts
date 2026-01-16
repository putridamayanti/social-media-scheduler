import {NextRequest, NextResponse} from "next/server";

export async function POST(req: NextRequest) {
    const res = await fetch(`${process.env.API_URL}/login`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Cookie: req.headers.get("cookie") || "",
        },
        body: await req.text(),
    });

    const body = await res.text();
    console.log(body)

    const response = new NextResponse(body, {
        status: res.status,
    });
    if (res.status === 200) {
        const setCookie = res.headers.get("set-cookie");
        if (setCookie) {
            response.headers.set("set-cookie", setCookie);
        }
    }

    return response;
}
