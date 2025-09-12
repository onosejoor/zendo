export function showDeleteBtn(role: TeamRole, memberRole: TeamRole) {
  if (role !== "owner" || memberRole === "owner") {
    return false;
  }
  return true;
}

export function checkRolesMatch(role: TeamRole, matcher: string[]) {
  return matcher.includes(role);
}
