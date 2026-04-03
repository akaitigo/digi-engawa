import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import type { LearnerProgress } from "@/types/progress";
import { ProgressOverview } from "./ProgressOverview";

const progressList: LearnerProgress[] = [
	{
		id: "prog-1",
		participant_id: "p-1",
		material_id: "m-1",
		current_step: 3,
		completed: false,
		updated_at: "2026-01-01T00:00:00Z",
	},
	{
		id: "prog-2",
		participant_id: "p-2",
		material_id: "m-1",
		current_step: 5,
		completed: true,
		updated_at: "2026-01-01T00:00:00Z",
	},
];

const participantNames: Record<string, string> = {
	"p-1": "田中太郎",
	"p-2": "山田花子",
};

describe("ProgressOverview", () => {
	it("renders participant names", () => {
		render(<ProgressOverview progressList={progressList} participantNames={participantNames} totalSteps={5} />);
		expect(screen.getByText("田中太郎")).toBeDefined();
		expect(screen.getByText("山田花子")).toBeDefined();
	});

	it("shows step progress for incomplete", () => {
		render(<ProgressOverview progressList={progressList} participantNames={participantNames} totalSteps={5} />);
		expect(screen.getByText("3/5 ステップ")).toBeDefined();
	});

	it("shows completed status", () => {
		render(<ProgressOverview progressList={progressList} participantNames={participantNames} totalSteps={5} />);
		expect(screen.getByText("完了")).toBeDefined();
	});

	it("shows empty message", () => {
		render(<ProgressOverview progressList={[]} participantNames={{}} totalSteps={5} />);
		expect(screen.getByText("進捗データはありません")).toBeDefined();
	});

	it("has accessible progressbar roles", () => {
		render(<ProgressOverview progressList={progressList} participantNames={participantNames} totalSteps={5} />);
		const bars = screen.getAllByRole("progressbar");
		expect(bars.length).toBe(2);
	});
});
