import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import type { Participant } from "@/types/classroom";
import { ParticipantList } from "./ParticipantList";

const participants: Participant[] = [
	{ id: "p-1", classroom_id: "c-1", name: "田中太郎", role: "learner", created_at: "2026-01-01T00:00:00Z" },
	{ id: "p-2", classroom_id: "c-1", name: "山田花子", role: "supporter", created_at: "2026-01-01T00:00:00Z" },
	{ id: "p-3", classroom_id: "c-1", name: "佐藤一郎", role: "organizer", created_at: "2026-01-01T00:00:00Z" },
];

describe("ParticipantList", () => {
	it("renders participant names", () => {
		render(<ParticipantList participants={participants} />);
		expect(screen.getByText("田中太郎")).toBeDefined();
		expect(screen.getByText("山田花子")).toBeDefined();
		expect(screen.getByText("佐藤一郎")).toBeDefined();
	});

	it("renders role labels", () => {
		render(<ParticipantList participants={participants} />);
		expect(screen.getByText("受講者")).toBeDefined();
		expect(screen.getByText("サポーター")).toBeDefined();
		expect(screen.getByText("主催者")).toBeDefined();
	});

	it("shows empty message", () => {
		render(<ParticipantList participants={[]} />);
		expect(screen.getByText("参加者はまだいません")).toBeDefined();
	});
});
