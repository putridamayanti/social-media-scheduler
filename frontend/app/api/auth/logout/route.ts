import {NextRequest, NextResponse} from "next/server";

export async function DELETE(req: NextRequest) {
    const res = await fetch(`${process.env.API_URL}/logout`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            Cookie: req.headers.get("cookie") || "",
        },
        body: await req.text(),
    });

    const body = await res.text();
    console.log(body)

    return new NextResponse(body, {
        status: res.status,
    });
}
