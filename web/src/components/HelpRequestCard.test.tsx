import { fireEvent, render, screen } from "@testing-library/react";
import { describe, expect, it, vi } from "vitest";
import type { HelpRequest } from "@/types/help";
import { HelpRequestCard } from "./HelpRequestCard";

const pendingRequest: HelpRequest = {
	id: "hr-1",
	classroom_id: "class-1",
	participant_id: "part-1",
	material_step_id: "step-1",
	status: "pending",
	created_at: "2026-01-01T10:00:00Z",
};

const inProgressRequest: HelpRequest = {
	...pendingRequest,
	status: "in_progress",
};

const resolvedRequest: HelpRequest = {
	...pendingRequest,
	status: "resolved",
	resolved_at: "2026-01-01T10:05:00Z",
};

describe("HelpRequestCard", () => {
	it("shows pending status", () => {
		render(<HelpRequestCard request={pendingRequest} />);
		expect(screen.getByText("🔴 たすけて")).toBeDefined();
	});

	it("shows action button for pending", () => {
		render(<HelpRequestCard request={pendingRequest} />);
		expect(screen.getByText("たいおうする")).toBeDefined();
	});

	it("calls onStatusChange when action button clicked", () => {
		const onStatusChange = vi.fn();
		render(<HelpRequestCard request={pendingRequest} onStatusChange={onStatusChange} />);
		fireEvent.click(screen.getByText("たいおうする"));
		expect(onStatusChange).toHaveBeenCalledWith("hr-1", "in_progress");
	});

	it("shows resolve button for in_progress", () => {
		render(<HelpRequestCard request={inProgressRequest} />);
		expect(screen.getByText("かいけつ")).toBeDefined();
	});

	it("shows no action button for resolved", () => {
		render(<HelpRequestCard request={resolvedRequest} />);
		expect(screen.getByText("🟢 かいけつ")).toBeDefined();
		expect(screen.queryByRole("button")).toBeNull();
	});
});
