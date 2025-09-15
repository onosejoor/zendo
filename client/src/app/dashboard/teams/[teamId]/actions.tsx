export function showDeleteBtn(
  memberId: string,
  userId: string,
  myRole: TeamRole
) {
  if (memberId === userId && myRole !== "owner") return true;

  if (myRole === "owner" && memberId !== userId) return true;

  return false;
}


export function checkRolesMatch(role: TeamRole, matcher: string[]) {
  return matcher.includes(role);
}
